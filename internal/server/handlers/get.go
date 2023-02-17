package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func (h *handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMetric")

	metric, exist := h.metricsStorage.Get(chi.URLParam(r, "name"))

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(metric.ToString()))
}
