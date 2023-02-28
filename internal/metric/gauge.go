package metric

import (
	"encoding/json"
	"github.com/lastbyte32/go-metric/internal/utils"
)

type gauge struct {
	name      string
	valueType MType
	value     float64
}

func (g *gauge) GetName() string {
	return g.name
}

func (g *gauge) GetType() MType {
	return g.valueType
}

func (g *gauge) ToString() string {
	return utils.FloatToString(g.value)
}

func (g *gauge) ToJSON() ([]byte, error) {
	m := Metrics{ID: g.name, MType: string(GAUGE), Value: &g.value}
	return json.Marshal(m)
}

func (g *gauge) SetValue(value string) error {
	f, err := utils.StringToFloat64(value)
	if err != nil {
		return err
	}
	g.value = f
	return nil
}

func (g *gauge) MarshalJSON() ([]byte, error) {
	return json.Marshal(Metrics{
		ID:    g.name,
		MType: string(GAUGE),
		Value: &g.value,
	})
}

func NewGauge(name string, value float64) IMetric {
	return &gauge{
		name,
		GAUGE,
		value,
	}
}
