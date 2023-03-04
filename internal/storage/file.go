package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/metric"
	"github.com/lastbyte32/go-metric/internal/utils"
)

const (
	WRITE = os.O_WRONLY | os.O_CREATE
	READ  = os.O_RDONLY
	PERM  = 0644
)

type fileStorage struct {
	*memoryStorage
	file       string
	interval   time.Duration
	isRestore  bool
	fileMutex  sync.RWMutex
	logger     *zap.SugaredLogger
	hash       string
	stopWorker chan int
}

func (store *fileStorage) Close() error {
	store.stopWorker <- 1
	close(store.stopWorker)
	err := store.saveInFile()
	if err != nil {
		return err
	}
	return nil
}

func (store *fileStorage) openFile(mode int) (*os.File, error) {
	file, err := os.OpenFile(store.file, mode, PERM)
	if err != nil {
		store.logger.Infof("err open file: %s, [%s]", store.file, err.Error())
		return nil, err
	}
	return file, err
}

func (store *fileStorage) storeWorkerOnInterval() {
	store.logger.Infof("store worker start, interval: %.0fs", store.interval.Seconds())
	go func() {
		ticker := time.NewTicker(store.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				store.saveInFile()
			case <-store.stopWorker:
				store.logger.Info("store worker stop")
				return
			}
		}
	}()
}

func (store *fileStorage) saveInFile() error {
	store.logger.Info("start store metric")
	metrics := store.All()
	if len(metrics) == 0 {
		store.logger.Info("metric empty, save skip")
		return nil
	}

	data, err := json.Marshal(metrics)
	if err != nil {
		store.logger.Infof("error  in JSON marshal. [%s]", err)
		return err
	}

	if !store.hasChanges(data) {
		store.logger.Info("no changes, skip save job")
		return nil
	}

	store.logger.Info(
		zap.String("metrics", string(data)),
	)

	store.fileMutex.Lock()
	defer store.fileMutex.Unlock()
	file, err := store.openFile(WRITE)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		store.logger.Infof("err write file. [%s]", err.Error())
		return err
	}

	store.logger.Info("saved JSON to file")
	return nil
}

func (store *fileStorage) restore() {
	store.logger.Infof("restore from file: %s", store.file)
	store.fileMutex.RLock()
	defer store.fileMutex.RUnlock()
	file, err := store.openFile(READ)
	if err != nil {
		return
	}
	defer file.Close()

	metricsFromFile := make(map[string]metric.Metrics)
	err = json.NewDecoder(file).Decode(&metricsFromFile)
	if err != nil {
		store.logger.Infof("file decode err: %s", err)
		return
	}

	for _, item := range metricsFromFile {
		err := store.Update(item.ID, item.GetValueAsString(), metric.MType(item.MType))
		if err != nil {
			store.logger.Infof("update %s", err.Error())
			return
		}
	}
}

func (store *fileStorage) getHash(data []byte) string {
	return utils.GetMD5Hash(data)
}

func (store *fileStorage) hasChanges(data []byte) bool {
	hash := store.getHash(data)
	if hash == store.hash {
		return false
	}
	store.hash = hash
	return true
}

func NewFileStorage(l *zap.SugaredLogger, config config.Configurator) IStorage {
	l.Info("new file storage")
	channel := make(chan int)
	store := &fileStorage{
		memoryStorage: &memoryStorage{
			logger: l,
			values: make(map[string]metric.IMetric),
		},
		logger:     l,
		interval:   config.GetStoreInterval(),
		file:       config.GetStoreFile(),
		isRestore:  config.IsRestore(),
		stopWorker: channel,
	}

	if store.isRestore {
		store.restore()
	}

	if store.interval != 0 {
		store.storeWorkerOnInterval()
	}

	return store
}
