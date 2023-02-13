package handlers

import (
	"github.com/lastbyte32/go-metric/internal/server/storage"
)

type Main struct {
	GaugeStorage    storage.GaugeStorage
	CountersStorage storage.CounterStorage
}
