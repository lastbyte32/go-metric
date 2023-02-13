package storage

import "fmt"

type CounterStorage interface {
	Update(string, int64)
	Get(name string) (int64, bool)
	All() map[string]int64
}

type CounterMemStorage struct {
	values map[string]int64
}

func (c *CounterMemStorage) All() map[string]int64 {
	return c.values
}

func (c *CounterMemStorage) Get(name string) (int64, bool) {
	fmt.Printf("get counter: %s\n", name)
	value, exist := c.values[name]
	return value, exist
}

func (c *CounterMemStorage) Update(name string, value int64) {
	c.values[name] += value
}

func NewCounterMemStorage() CounterStorage {
	return &CounterMemStorage{map[string]int64{}}
}
