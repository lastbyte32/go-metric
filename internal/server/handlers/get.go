package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/metric"
	"net/http"
)

func (h *handler) GetMetricFromJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inputMetric metric.Metrics
	if err := json.NewDecoder(r.Body).Decode(&inputMetric); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mtype := metric.MType(inputMetric.MType)
	if mtype != metric.COUNTER && mtype != metric.GAUGE {
		fmt.Println("invalid_type")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"error": "err invalid_type"}`))
		return
	}

	metric, exist := h.metricsStorage.Get(inputMetric.ID)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "metric not found"}`))
		return
	}
	json, err := metric.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}
	w.Write(json)
}

func (h *handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("GetMetric")
	mType := metric.MType(chi.URLParam(r, "type"))
	if mType != metric.COUNTER && mType != metric.GAUGE {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	metricName := chi.URLParam(r, "name")
	m, exist := h.metricsStorage.Get(metricName)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("metric not found"))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(m.ToString()))
}
