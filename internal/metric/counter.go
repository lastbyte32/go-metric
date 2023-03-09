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
	hash      string
}

func (c *counter) SetHash(key string) error {
	part := fmt.Sprintf("%s:%s:%d", c.name, c.valueType, c.value)
	hash, err := utils.GetSha256Hash(part, key)
	if err != nil {
		return err
	}
	c.hash = hash
	return nil
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

func (c *counter) SetValue(value string) error {
	s, err := utils.StringToInt64(value)
	if err != nil {
		fmt.Println("COUNTER err parse")
		return err
	}
	c.value += s
	return nil
}

func (c *counter) MarshalJSON() ([]byte, error) {
	return json.Marshal(Metrics{
		ID:    c.name,
		MType: string(COUNTER),
		Delta: c.value,
		Hash:  c.hash,
	})
}

func NewCounter(name string, value int64) IMetric {
	return &counter{
		name:      name,
		valueType: COUNTER,
		value:     value,
	}
}
