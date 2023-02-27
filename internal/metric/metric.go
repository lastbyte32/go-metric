package metric

import (
	"errors"
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
}

type IMetric interface {
	GetName() string
	GetType() MType
	ToString() string
	ToJson() ([]byte, error)
	SetValue(value string) error
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
