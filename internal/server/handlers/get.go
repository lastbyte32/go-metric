package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"net/http"
)

func (h *handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMetric")

	name := chi.URLParam(r, "name")
	found := false
	value := ""

	switch storage.MType(chi.URLParam(r, "type")) {
	case storage.GAUGE:
		metric, exist := h.metricsStorage.Get(name)
		if exist {
			value = metric.GetGauge()
			found = true
		}
	case storage.COUNTER:
		metric, exist := h.metricsStorage.Get(name)
		if exist {
			value = metric.GetCounter()
			found = true
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	if !found {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write([]byte(value))
}
