package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func response(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(body))
	w.WriteHeader(status)
}

func (h *Main) GetOneMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetOneMetric")

	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")

	switch metric {
	case "gauge":
		value, exist := h.GaugeStorage.Get(name)
		if exist {
			response(w, http.StatusOK, fmt.Sprintf("%v", value))
		} else {
			response(w, http.StatusNotFound, fmt.Sprintf("metric name: %s not found", name))
		}
		return
	case "counter":
		value, exist := h.CountersStorage.Get(name)
		if exist {
			response(w, http.StatusOK, fmt.Sprintf("%v", value))
		} else {
			response(w, http.StatusNotFound, fmt.Sprintf("metric name: %s not found", name))
		}
	default:
		fmt.Printf("wrong metric type: %v\n", metric)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
