package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/metric"
	"net/http"
	"strconv"
)

func (h *handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateHandle")

	valueType := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")
	value := chi.URLParam(r, "value")

	switch metric.MType(valueType) {
	case metric.GAUGE:
		fmt.Println("case gauge")
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("err parse value"))
			return
		}
		h.metricsStorage.Update(name, value, metric.GAUGE)

	case metric.COUNTER:
		fmt.Println("case counter")
		_, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("err parse value"))
			return
		}
		h.metricsStorage.Update(name, value, metric.COUNTER)

	default:
		fmt.Println("invalid_type")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("err invalid_type"))
	}
}
