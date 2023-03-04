package handlers

import (
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/storage"
)

type handler struct {
	metricsStorage storage.IStorage
	logger         *zap.SugaredLogger
}

func NewHandler(s storage.IStorage, l *zap.SugaredLogger) *handler {
	return &handler{
		metricsStorage: s,
		logger:         l,
	}
}
