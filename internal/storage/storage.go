// Package storage Модуль хранения данных
package storage

import (
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/metric"
)

type Type int

const (
	MEMORY Type = 1 << iota
	FILE
	SQL
)

type IStorage interface {
	// Get - получить метрику по строковому имени
	Get(name string) (metric.IMetric, bool)
	// All - получить все метрики
	All() map[string]metric.IMetric
	// Update - обновить метрику
	Update(name string, value string, metricType metric.MType) error
	// Close - безопасно закрыть хранилище
	Close() error
}

// New - конструктор создает хранилище определенного типа указанного в конфигурации
func New(config config.Configurator, logger *zap.SugaredLogger) IStorage {
	switch getStorageType(config) {
	case MEMORY:
		return NewMemoryStorage(logger)
	case FILE:
		return NewFileStorage(logger, config)
	case SQL:
		return NewSQLStorage(logger, config)
	default:
		return NewMemoryStorage(logger)
	}
}

func getStorageType(config config.Configurator) Type {
	if config.GetDSN() != "" {
		return SQL
	}
	if config.GetStoreFile() != "" {
		return FILE
	}
	return MEMORY
}
