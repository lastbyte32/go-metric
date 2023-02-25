package storage

import (
	"github.com/lastbyte32/go-metric/internal/metric"
)

type Storage interface {
	Get(name string) (metric.Metric, bool)
	All() map[string]metric.Metric
	Update(name string, value string, metricType metric.MType) error
}
