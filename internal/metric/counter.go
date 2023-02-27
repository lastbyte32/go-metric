package metric

import (
	"encoding/json"
	"fmt"
	"github.com/lastbyte32/go-metric/internal/utils"
)

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

func (c *counter) ToJson() ([]byte, error) {
	m := Metrics{ID: c.name, MType: string(COUNTER), Delta: &c.value}
	return json.Marshal(m)
}

func (c *counter) SetValue(value string) error {
	//fmt.Println("Increase")
	s, err := utils.StringToInt64(value)
	if err != nil {
		fmt.Println("COUNTER err parse")
		return err
	}
	c.value += s
	return nil
}

func NewCounter(name string, value int64) IMetric {
	return &counter{
		name,
		COUNTER,
		value,
	}
}
