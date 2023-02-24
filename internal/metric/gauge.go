package metric

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/utils"
)

type gauge struct {
	name  string
	value float64
}

func (g *gauge) GetName() string {
	return g.name
}

func (g *gauge) GetType() MType {
	return GAUGE
}

func (g *gauge) ToString() string {
	return fmt.Sprintf("%.3f", g.value)
}

func (g *gauge) SetValue(value string) error {
	f, err := utils.StringToFloat64(value)
	if err != nil {
		return err
	}
	g.value = f
	return nil
}

func NewGauge(name string, value float64) Metric {
	return &gauge{
		name,
		value,
	}
}
