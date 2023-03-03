package storage

import (
	"sync"

	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/metric"
)

type memoryStorage struct {
	values map[string]metric.IMetric
	logger *zap.SugaredLogger
	mutex  sync.RWMutex
}

func (ms *memoryStorage) Close() error {
	// хранилище in-memory, нечего делать при завершении
	return nil
}

func (ms *memoryStorage) Get(name string) (metric.IMetric, bool) {
	ms.logger.Infof("store: get %s", name)
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	oneMetric, ok := ms.values[name]
	if !ok {
		return nil, false
	}
	return oneMetric, true
}

func (ms *memoryStorage) Update(name, value string, metricType metric.MType) error {
	ms.logger.Infof("store: update %s -> %s", name, metricType)

	m, found := ms.Get(name)
	if !found {
		newMetric, err := metric.NewByString(name, value, metricType)
		if err != nil {
			return err
		}
		ms.mutex.Lock()
		ms.values[name] = newMetric
		ms.mutex.Unlock()
		return nil
	}
	ms.mutex.Lock()
	err := m.SetValue(value)
	ms.mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func (ms *memoryStorage) All() map[string]metric.IMetric {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	return ms.values
}

func NewMemoryStorage(l *zap.SugaredLogger) IStorage {
	l.Info("new memory storage")
	return &memoryStorage{
		values: make(map[string]metric.IMetric),
		logger: l,
	}
}
