package handlers

import (
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/storage"
)

type handler struct {
	metricsStorage storage.IStorage
	logger         *zap.SugaredLogger
	config         config.Configurator
}

func NewHandler(s storage.IStorage, l *zap.SugaredLogger, c config.Configurator) *handler {
	return &handler{
		metricsStorage: s,
		logger:         l,
		config:         c,
	}
}
