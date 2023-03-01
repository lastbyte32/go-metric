package handlers

import (
	"github.com/lastbyte32/go-metric/internal/storage"
	"go.uber.org/zap"
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
