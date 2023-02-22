package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/metric"
	"net/http"
)

func (h *handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMetric")
	mType := metric.MType(chi.URLParam(r, "type"))

	if mType != metric.COUNTER && mType != metric.GAUGE {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	m, exist := h.metricsStorage.Get(chi.URLParam(r, "name"), mType)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(m.ToString()))
}
