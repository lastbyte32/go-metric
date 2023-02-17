package storage

type GaugeStorage interface {
	Update(string, float64)
	Get(name string) (float64, bool)
	All() map[string]float64
}

type GaugeMemStorage struct {
	values map[string]float64
}

func (g *GaugeMemStorage) All() map[string]float64 {
	return g.values
}

func (g *GaugeMemStorage) Get(name string) (float64, bool) {
	value, exist := g.values[name]
	return value, exist
}

func (g *GaugeMemStorage) Update(name string, value float64) {
	g.values[name] = value
}

func NewGaugeMemStorage() GaugeStorage {
	return &GaugeMemStorage{map[string]float64{}}
}
