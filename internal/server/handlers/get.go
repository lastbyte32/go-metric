package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"net/http"
)

func response(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(body))
}

func (h *handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMetric")

	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")

	switch storage.MType(metric) {
	case storage.GAUGE:
		metric, exist := h.metricsStorage.Get(name)
		if !exist {
			response(w, http.StatusNotFound, fmt.Sprintf("metric name: %s not found", name))
			return
		}
		response(w, http.StatusOK, fmt.Sprintf("%v", metric.GetGauge()))
		return
	case storage.COUNTER:
		metric, exist := h.metricsStorage.Get(name)
		if exist {
			response(w, http.StatusNotFound, fmt.Sprintf("metric name: %s not found", name))
			return
		}
		response(w, http.StatusOK, fmt.Sprintf("%v", metric.GetCounter()))
		return

	default:
		fmt.Printf("wrong metric type: %v\n", metric)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
