package storage

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/metric"
	"strconv"
	"sync"
)

type memoryStorage struct {
	counter map[string]metric.Metric
	gauge   map[string]metric.Metric
	sync.Mutex
}

func (ms *memoryStorage) Get(name string, valueType metric.MType) (metric.Metric, bool) {
	ms.Lock()
	defer ms.Unlock()

	switch valueType {
	case metric.GAUGE:
		fmt.Println("Get GAUGE")
		m, ok := ms.gauge[name]
		if ok {
			return m, true
		}
		fmt.Printf("Get GAUGE [%s]: not found", name)
		return nil, false

	case metric.COUNTER:
		m, ok := ms.counter[name]
		if ok {
			return m, true
		}
		fmt.Printf("Get COUNTER [%s]: not found", name)
		return nil, false
	default:
		return nil, false
	}

}

func (ms *memoryStorage) All() map[string]metric.Metric {
	ms.Lock()
	defer ms.Unlock()
	all := map[string]metric.Metric{}
	for k, v := range ms.counter {
		all[k] = v
	}
	for k, v := range ms.gauge {
		all[k] = v
	}

	return all
}

func (ms *memoryStorage) Update(name, value string, valueType metric.MType) {
	ms.Lock()
	defer ms.Unlock()

	switch valueType {
	case metric.GAUGE:
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			fmt.Println("GAUGE update")
			ms.gauge[name] = metric.NewGauge(name, f)
		} else {
			fmt.Println("GAUGE err parse")
		}
	case metric.COUNTER:
		existMetric, ok := ms.counter[name]
		if ok {
			s, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				existMetric.Increase(s)
			}
		}
	}

}
func NewMemoryStorage() Storage {
	return &memoryStorage{
		gauge:   map[string]metric.Metric{},
		counter: map[string]metric.Metric{},
	}
}
