package storage

import (
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/metric"
)

type Type int

type IStorage interface {
	Get(name string) (metric.IMetric, bool)
	All() map[string]metric.IMetric
	Update(name string, value string, metricType metric.MType) error
	Close() error
}

// TODO удалю этот код после ревью
//const (
//	MEMORY Type = iota
//	FILE
//)
//func New(config config.Configurator, logger *zap.SugaredLogger) IStorage {
//
//	switch config.GetStorageType() {
//	case MEMORY:
//		return NewMemoryStorage(logger)
//	case FILE:
//		return NewFileStorage(logger, config)
//	default:
//		return NewMemoryStorage(logger)
//	}
//}

func New(config config.Configurator, logger *zap.SugaredLogger) IStorage {
	if config.GetStoreFile() != "" {
		return NewFileStorage(logger, config)
	}
	return NewMemoryStorage(logger)
}
