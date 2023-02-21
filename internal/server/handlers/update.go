package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"net/http"
	"strconv"
)

func (h *handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateHandle")

	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")
	value := chi.URLParam(r, "value")

	switch storage.MType(metric) {
	case storage.GAUGE:
		fmt.Println("case gauge")
		valueGauge, err := strconv.ParseFloat(value, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("err parse value"))
			return
		}
		h.metricsStorage.Update(name, storage.NewMetric(
			name,
			storage.GAUGE,
			valueGauge,
			0,
		))

	case storage.COUNTER:
		fmt.Println("case counter")
		valueCounter, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("err parse value"))
			return
		}
		h.metricsStorage.Update(name, storage.NewMetric(
			name,
			storage.COUNTER,
			0,
			valueCounter,
		))
	default:
		fmt.Println("invalid_type")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("err invalid_type"))
	}

}
