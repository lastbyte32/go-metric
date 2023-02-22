package metric

import "fmt"

type Gauge struct {
	name      string
	valueType MType
	value     float64
}

func (g *Gauge) Increase(int642 int64) {
	//TODO implement me
	panic("implement me")
}

func (g *Gauge) GetName() string {
	return g.name
}

func (g *Gauge) GetType() MType {
	return g.valueType
}

func (g *Gauge) ToString() string {
	return fmt.Sprintf("%.3f", g.value)
}

func (g *Gauge) SetValue(value float64) {
	g.value = value
}

func NewGauge(name string, value float64) Metric {
	return &Gauge{name: name, valueType: GAUGE, value: value}
}
