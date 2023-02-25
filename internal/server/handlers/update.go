package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/metric"
	"net/http"
)

func (h *handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateHandle")
	metricType := metric.MType(chi.URLParam(r, "type"))
	metricName := chi.URLParam(r, "name")

	if metricType != metric.COUNTER && metricType != metric.GAUGE {
		fmt.Println("invalid_type")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("err invalid_type"))
		return
	}

	value := chi.URLParam(r, "value")
	err := h.metricsStorage.Update(metricName, value, metricType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}
