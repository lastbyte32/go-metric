package storage

import "fmt"

type MType string

const (
	COUNTER MType = "counter"
	GAUGE   MType = "gauge"
)

type Storage interface {
	Get(name string) (Metric, bool)
	All() map[string]Metric
	Update(string, Metric)
}

type Metric struct {
	name      string
	valueType MType
	gauge     float64
	counter   int64
}

func (m *Metric) GetGauge() string {
	return fmt.Sprintf("%f", m.gauge)
}

func (m *Metric) GetCounter() string {
	return fmt.Sprintf("%d", m.counter)
}

func (m *Metric) GetName() string {
	return m.name
}

func (m *Metric) ToString() string {
	switch m.valueType {
	case GAUGE:
		return m.GetGauge()
	case COUNTER:
		return m.GetCounter()
	default:
		return ""
	}
}

func NewMetric(name string, valueType MType, gauge float64, counter int64) Metric {
	return Metric{
		name,
		valueType,
		gauge,
		counter,
	}
}

type memoryStorage struct {
	values map[string]Metric
}

func (m *memoryStorage) Get(name string) (Metric, bool) {
	metric, ok := m.values[name]
	if !ok {
		return Metric{}, false
	}
	return metric, true
}

func (m *memoryStorage) All() map[string]Metric {
	return m.values
}

func (m *memoryStorage) Update(name string, metric Metric) {

	// TODO сделать метод потокобезопасным
	// TODO: кмк тут слишком сложно,
	// Возможно при записи метрики нужно проверять не отличается ли тип,
	// если различается, то ругаться и не перезаписывать
	switch metric.valueType {
	case GAUGE:
		m.values[name] = metric
	case COUNTER:
		existMetric, ok := m.Get(name)
		if ok {
			existMetric.counter += metric.counter
			m.values[name] = existMetric
		} else {
			m.values[name] = metric
		}
	}

}
func NewMemoryStorage() Storage {
	return &memoryStorage{map[string]Metric{}}
}
