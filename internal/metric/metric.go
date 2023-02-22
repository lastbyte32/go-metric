package metric

import "fmt"

type MType string

const (
	COUNTER MType = "counter"
	GAUGE   MType = "gauge"
)

type Metric interface {
	GetName() string
	GetType() MType
	ToString() string
}

type MemoryMetric struct {
	name      string
	valueType MType
	gauge     float64
	counter   int64
}

func (m *MemoryMetric) GetName() string {
	return m.name
}

func (m *MemoryMetric) GetType() MType {
	return m.valueType
}

func (m *MemoryMetric) ToString() string {
	switch m.valueType {
	case GAUGE:
		return fmt.Sprintf("%.3f", m.gauge)
	case COUNTER:
		return fmt.Sprintf("%d", m.counter)
	default:
		return ""
	}
}

func NewCounter(name string, value int64) Metric {
	return &MemoryMetric{name: name, valueType: COUNTER, counter: value}
}

func NewGauge(name string, value float64) Metric {
	return &MemoryMetric{name: name, valueType: GAUGE, gauge: value}
}

func (m *MemoryMetric) Increase(name, string, value int64) {
	m.counter += value
}
