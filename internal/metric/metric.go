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

type Metric interface {
	GetName() string
	GetType() MType
	ToString() string
	SetValue(value string) error
}

func NewByString(name string, value string, metricType MType) (error, Metric) {
	switch metricType {

	case GAUGE:
		err, f := utils.StringToFloat64(value)
		if err != nil {
			return err, nil
		}
		return nil, &gauge{name: name, valueType: GAUGE, value: f}
	case COUNTER:
		err, f := utils.StringToInt64(value)
		if err != nil {
			return err, nil
		}
		return nil, &counter{name: name, valueType: COUNTER, value: f}

	default:
		return errors.New("wrong metric type"), nil
	}
}
