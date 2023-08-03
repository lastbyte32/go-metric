package metric

import (
	"errors"
	"fmt"

	mproto "github.com/lastbyte32/go-metric/api/proto"
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

func (m *Metrics) VerifyHash(key string) (bool, error) {
	hash, err := utils.GetSha256Hash(m.GetStringToSign(), key)
	if err != nil {
		return false, err
	}
	if m.Hash != hash {
		return false, nil
	}
	return true, nil
}

// IMetric - интерфейс метрики.
type IMetric interface {
	// GetName - получить имя метрики.
	GetName() string
	// GetType - получить тип метрики.
	GetType() MType
	// ToString - получить значение в виде строки.
	ToString() string
	MarshalJSON() ([]byte, error)
	MarshalProtobuf() *mproto.Metric
	// SetValue - установить значение.
	SetValue(value string) error
	// SetHash - создает хеш в структуру метрики используя ключ.
	SetHash(key string) error
}

// NewByString - конструктор создает метрику из строки.
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
