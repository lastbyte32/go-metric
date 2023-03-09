package metric

import (
	"errors"
	"fmt"

	"github.com/lastbyte32/go-metric/internal/utils"
)

type MType string

const (
	COUNTER MType = "counter"
	GAUGE   MType = "gauge"
)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	Hash  string   `json:"hash,omitempty"`  // значение хеш-функции
}

func (m *Metrics) GetValueAsString() string {
	// Или этот метод нужен только в рамках файлового хранилища и его нужно имплементировать там?
	switch MType(m.MType) {
	case COUNTER:
		return fmt.Sprintf("%d", *m.Delta)
	case GAUGE:
		return utils.FloatToString(*m.Value)
	default:
		return ""
	}
}

func (m *Metrics) GetStringToSign() string {
	switch MType(m.MType) {
	case COUNTER:
		return fmt.Sprintf("%s:%s:%d", m.ID, m.MType, *m.Delta)
	case GAUGE:
		return fmt.Sprintf("%s:%s:%f", m.ID, m.MType, *m.Value)
	default:
		return ""
	}
}

type IMetric interface {
	GetName() string
	GetType() MType
	ToString() string
	MarshalJSON() ([]byte, error)
	SetValue(value string) error
	SetHash(key string) error
}

func NewByString(name string, value string, metricType MType) (IMetric, error) {
	switch metricType {
	case GAUGE:
		f, err := utils.StringToFloat64(value)
		if err != nil {
			return nil, err
		}
		return &gauge{name: name, valueType: GAUGE, value: f}, nil
	case COUNTER:
		f, err := utils.StringToInt64(value)
		if err != nil {
			return nil, err
		}
		return &counter{name: name, valueType: COUNTER, value: f}, nil

	default:
		return nil, errors.New("NewByString: wrong metric type")
	}
}
