package storage

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/metric"
	"sync"
)

type memoryStorage struct {
	values map[string]metric.Metric
	sync.Mutex
}

func (ms *memoryStorage) Get(name string) (metric.Metric, bool) {
	fmt.Println("get metric")
	ms.Lock()
	defer ms.Unlock()
	m, ok := ms.values[name]
	if ok {
		fmt.Printf("Get metric [%s]: OK", name)
		return m, true
	}
	fmt.Printf("Get metric [%s]: not found", name)
	return nil, false
}

func (ms *memoryStorage) Update(name, value string, metricType metric.MType) error {
	fmt.Println("ms->update")

	defer ms.Unlock()

	m, found := ms.Get(name)
	if !found {
		fmt.Println("create new metric")
		newMetric, err := metric.NewByString(name, value, metricType)
		if err != nil {
			return err
		}
		fmt.Println("update ok")
		ms.Lock()
		ms.values[name] = newMetric
		return nil
	}
	ms.Lock()
	err := m.SetValue(value)
	if err != nil {
		return err
	}
	return nil
}

func (ms *memoryStorage) All() map[string]metric.Metric {
	ms.Lock()
	defer ms.Unlock()
	return ms.values
}

func NewMemoryStorage() Storage {
	fmt.Println("new memory storage")
	return &memoryStorage{
		values: make(map[string]metric.Metric),
	}
}
