package metric

import "fmt"

type counter struct {
	name      string
	valueType MType
	value     int64
}

func (c *counter) GetName() string {
	return c.name
}

func (c *counter) GetType() MType {
	return c.valueType
}

func (c *counter) ToString() string {
	return fmt.Sprintf("%d", c.value)
}

func (c *counter) Increase(value int64) {
	c.value += value
}

func NewCounter(name string, value int64) Metric {
	return &counter{name: name, valueType: COUNTER, value: value}
}
