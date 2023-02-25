package handlers

import (
	"github.com/lastbyte32/go-metric/internal/server/storage"
)

type handler struct {
	metricsStorage storage.Storage
}

func NewHandler(storage storage.Storage) *handler {
	return &handler{
		storage,
	}
}
