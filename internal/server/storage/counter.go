package storage

type CounterStorage interface {
	Update(string, int64)
}

type CounterMemStorage struct {
	values map[string]int64
}

func (c *CounterMemStorage) Update(name string, value int64) {
	c.values[name] += value
}

func NewCounterMemStorage() CounterStorage {
	return &CounterMemStorage{map[string]int64{}}
}
