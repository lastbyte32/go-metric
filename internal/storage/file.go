package storage

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/metric"
	"github.com/lastbyte32/go-metric/internal/utils"
)

type fileStorage struct {
	*memoryStorage

	file           string
	interval       time.Duration
	isRestore      bool
	eventCh        chan int
	fileMutex      sync.RWMutex
	logger         *zap.SugaredLogger
	hash           [16]byte
	shutdownSignal chan int
}

func (store *fileStorage) storeWorkerOnInterval(ctx context.Context) {
	store.logger.Infof("store worker start, interval: %.0fs", store.interval.Seconds())
	go func() {
		ticker := time.NewTicker(store.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				store.saveInFile()
			case <-ctx.Done():
				store.saveInFile()
				store.logger.Info("shutdown store worker")
				return
			}
		}
	}()
}

func (store *fileStorage) saveInFile() {
	store.logger.Info("start store metric")
	metrics := store.All()
	if len(metrics) == 0 {
		store.logger.Info("metric empty, save skip")
		return
	}
	data, err := json.Marshal(metrics)
	if err != nil {
		store.logger.Infof("error  in JSON marshal. [%s]", err)
		return
	}
	jsonHash := md5.Sum(data)
	if store.hash == jsonHash {
		store.logger.Info("hash equal, save skip")
		return
	}
	store.hash = jsonHash

	store.logger.Info(
		zap.String("metrics", string(data)),
	)

	store.fileMutex.Lock()
	defer store.fileMutex.Unlock()
	file, err := os.OpenFile(store.file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		store.logger.Infof("err open file: %s, [%s]", store.file, err.Error())
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		store.logger.Infof("err write file. [%s]", err.Error())
		return
	}

	store.logger.Info("saved JSON to file")
}

func (store *fileStorage) restore() {
	store.logger.Infof("restore from file: %s", store.file)
	store.fileMutex.RLock()
	defer store.fileMutex.RUnlock()
	file, err := os.OpenFile(store.file, os.O_RDONLY, 0777)
	if err != nil {
		store.logger.Infof("err open file: [%s]", err.Error())
		return
	}
	defer file.Close()

	metricsFromFile := make(map[string]metric.Metrics)
	err = json.NewDecoder(file).Decode(&metricsFromFile)
	if err != nil {
		store.logger.Infof("file decode err: %s", err)
	}
	value := ""
	for _, item := range metricsFromFile {
		switch metric.MType(item.MType) {
		case metric.COUNTER:
			value = fmt.Sprintf("%d", *item.Delta)
		case metric.GAUGE:
			value = utils.FloatToString(*item.Value)
		default:
			continue
		}

		if value == "" {
			continue
		}

		err := store.Update(item.ID, value, metric.MType(item.MType))
		if err != nil {
			store.logger.Infof("update %s", err.Error())
			return
		}
	}
}

func (store *fileStorage) Init(ctx context.Context) error {

	if store.isRestore {
		store.restore()
	}

	if store.interval == 0 {
		go func() {
			store.logger.Info("start shutdown watcher")
			<-ctx.Done()
			store.saveInFile()
			store.logger.Info("shutdown file storage")
		}()
	} else {
		store.storeWorkerOnInterval(ctx)
	}

	return nil
}

func NewFileStorage(l *zap.SugaredLogger, storeFile string, storeInterval time.Duration, isRestore bool) IStorage {
	l.Info("new file storage")

	store := &fileStorage{
		memoryStorage: &memoryStorage{
			logger: l,
			values: make(map[string]metric.IMetric),
		},
		interval:  storeInterval,
		file:      storeFile,
		logger:    l,
		isRestore: isRestore,
	}

	return store
}
