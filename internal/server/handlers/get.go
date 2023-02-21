package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"net/http"
)

func (h *handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMetric")
	mType := storage.MType(chi.URLParam(r, "type"))

	if mType != storage.GAUGE || mType != storage.COUNTER {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	metric, exist := h.metricsStorage.Get(chi.URLParam(r, "name"))

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(metric.ToString()))
}
