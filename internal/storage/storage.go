package storage

import (
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/metric"
)

type IStorage interface {
	Get(name string) (metric.IMetric, bool)
	All() map[string]metric.IMetric
	Update(name string, value string, metricType metric.MType) error
	Close() error
}

func New(config config.IStorage, logger *zap.SugaredLogger) IStorage {
	if config.GetStoreFile() != "" {
		return NewFileStorage(logger, config)
	}
	return NewMemoryStorage(logger)
}
