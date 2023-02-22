package storage

import (
	"github.com/lastbyte32/go-metric/internal/metric"
)

type Storage interface {
	Get(string, metric.MType) (metric.Metric, bool)
	All() map[string]metric.Metric
	Update(string, string, metric.MType)
}
