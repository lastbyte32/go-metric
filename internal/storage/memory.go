package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/lastbyte32/go-metric/internal/metric"
	"github.com/lastbyte32/go-metric/internal/utils"
)

type memoryStorage struct {
	values        map[string]metric.IMetric
	storeFile     string
	storeInterval time.Duration
	isRestore     bool
	ctx           context.Context
	eventCh       chan int
	fileMutex     sync.Mutex
	sync.Mutex
}

func (ms *memoryStorage) Get(name string) (metric.IMetric, bool) {
	//fmt.Println("get metric")
	ms.Lock()
	defer ms.Unlock()
	m, ok := ms.values[name]
	if ok {
		//fmt.Printf("Get metric [%s]: OK", name)
		return m, true
	}
	//fmt.Printf("Get metric [%s]: not found", name)
	return nil, false
}

func (ms *memoryStorage) Update(name, value string, metricType metric.MType) error {
	//fmt.Println("ms->update")
	defer ms.eventSaved()

	m, found := ms.Get(name)
	if !found {
		//fmt.Println("create new metric")
		newMetric, err := metric.NewByString(name, value, metricType)
		if err != nil {
			return err
		}
		//fmt.Println("update ok")
		ms.Lock()
		ms.values[name] = newMetric
		ms.Unlock()
		if ms.storeInterval == 0 && ms.storeFile != "" {
			ms.storeOnFile()
		}
		return nil
	}
	ms.Lock()
	err := m.SetValue(value)
	ms.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func (ms *memoryStorage) All() map[string]metric.IMetric {
	ms.Lock()
	defer ms.Unlock()
	return ms.values
}

func (ms *memoryStorage) eventSaved() {
	if ms.storeFile == "" {
		return
	}
	if ms.storeInterval != 0 {
		return
	}
	ms.eventCh <- 1
}

func (ms *memoryStorage) storeWorkerOnInterval(interval time.Duration) {
	fmt.Printf("store worker start, interval: %.0fs\n", interval.Seconds())
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ms.storeOnFile()
			case <-ms.ctx.Done():
				ms.storeOnFile()
				fmt.Println("store worker stop")
				return
			}
		}
	}()
}

func (ms *memoryStorage) storeWorkerOnUpdate() {
	fmt.Println("storeWorkerOnUpdate start")
	go func() {
		for {
			select {
			case <-ms.eventCh:
				ms.storeOnFile()
				fmt.Println("SAVED EVENT")
			case <-ms.ctx.Done():
				fmt.Println("storeWorkerOnUpdate stop")
				return
			}
		}
	}()
}

func (ms *memoryStorage) storeOnFile() {
	log.Println("start store metric")

	data, err := json.Marshal(ms.All())
	if err != nil {
		log.Printf("error  in JSON marshal. [%s]", err)
		return
	}
	log.Print(string(data))
	ms.fileMutex.Lock()
	defer ms.fileMutex.Unlock()

	file, err := os.OpenFile(ms.storeFile, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Printf("err open db: [%s]\n", err.Error())
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Printf("err write db file. [%s]", err.Error())
		return
	}

	log.Printf("saved JSON to file")
}

func WithStore(storeFile string, interval time.Duration) func(*memoryStorage) {
	return func(ms *memoryStorage) {
		fmt.Printf("WithStore on file: %s\n", storeFile)
		ms.storeFile = storeFile
		ms.storeInterval = interval
		if ms.storeInterval == 0 {
			ms.storeWorkerOnUpdate()
		} else {
			ms.storeWorkerOnInterval(interval)
		}

	}
}

func WithContext(ctx context.Context) func(*memoryStorage) {
	return func(ms *memoryStorage) {
		ms.ctx = ctx
	}
}

func WithRestore(fileName string, isRestore bool) func(*memoryStorage) {
	if !isRestore {
		fmt.Println("Restore skip")
		return func(*memoryStorage) {}
	}

	if fileName == "" {
		fmt.Println("Restore skip")
		return func(*memoryStorage) {}
	}

	return func(ms *memoryStorage) {
		ms.storeFile = fileName
		fmt.Printf("WithRestore from file: %s\n", fileName)
		ms.fileMutex.Lock()
		defer ms.fileMutex.Unlock()
		file, err := os.OpenFile(ms.storeFile, os.O_RDONLY, 0777)
		if err != nil {
			fmt.Printf("err open db: [%s]\n", err.Error())
		}
		defer file.Close()

		metricsFromFile := make(map[string]metric.Metrics)
		err = json.NewDecoder(file).Decode(&metricsFromFile)
		if err != nil {
			log.Println(err)
		}
		value := ""
		for _, item := range metricsFromFile {
			fmt.Println(item.ID)
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

			err := ms.Update(item.ID, value, metric.MType(item.MType))
			if err != nil {
				fmt.Printf("update %s", err.Error())
				return
			}
		}
	}
}

func NewMemoryStorage(options ...func(*memoryStorage)) IStorage {
	ms := &memoryStorage{
		eventCh:   make(chan int),
		isRestore: false,
		values:    make(map[string]metric.IMetric),
	}

	for _, o := range options {
		o(ms)
	}
	return ms
}
