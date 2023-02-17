package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
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

	switch metric {
	case "gauge":
		metric, exist := h.metricsStorage.Get(name)
		if exist {
			response(w, http.StatusOK, fmt.Sprintf("%v", metric.GetGauge()))
		} else {
			response(w, http.StatusNotFound, fmt.Sprintf("metric name: %s not found", name))
		}
		return
	case "counter":
		metric, exist := h.metricsStorage.Get(name)
		if exist {
			response(w, http.StatusOK, fmt.Sprintf("%v", metric.GetCounter()))
		} else {
			response(w, http.StatusNotFound, fmt.Sprintf("metric name: %s not found", name))
		}
	default:
		fmt.Printf("wrong metric type: %v\n", metric)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
