package handlers

import (
	"github.com/lastbyte32/go-metric/internal/storage"
)

type handler struct {
	metricsStorage storage.IStorage
}

func NewHandler(storage storage.IStorage) *handler {
	return &handler{
		storage,
	}
}
