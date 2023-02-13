package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *Main) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateHandle")

	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")
	value := chi.URLParam(r, "value")

	switch metric {
	case "gauge":
		fmt.Println("case gauge")
		valueGauge, err := strconv.ParseFloat(value, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("err parse value"))
			return
		}
		h.GaugeStorage.Update(name, valueGauge)

	case "counter":
		fmt.Println("case counter")
		valueCounter, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("err parse value"))
			return
		}
		h.CountersStorage.Update(name, valueCounter)
	default:
		fmt.Println("invalid_type")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("err invalid_type"))
	}

}
