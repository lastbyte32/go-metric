package storage

import (
	"github.com/lastbyte32/go-metric/internal/metric"
)

type IStorage interface {
	Get(name string) (metric.IMetric, bool)
	All() map[string]metric.IMetric
	Update(name string, value string, metricType metric.MType) error
}
