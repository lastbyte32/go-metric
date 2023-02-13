package storage

type GaugeStorage interface {
	Update(string, float64)
}

type GaugeMemStorage struct {
	values map[string]float64
}

func (g *GaugeMemStorage) Update(name string, value float64) {
	g.values[name] = value
}

func NewGaugeMemStorage() GaugeStorage {
	return &GaugeMemStorage{map[string]float64{}}
}
