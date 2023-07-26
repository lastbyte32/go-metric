package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/lastbyte32/go-metric/internal/metric"
)

// GetMetricFromJSON - отдает одну метрику запрощенную методом POST в виде JSON.
func (h *handler) GetMetricFromJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inputMetric metric.Metrics
	if err := json.NewDecoder(r.Body).Decode(&inputMetric); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mType := metric.MType(inputMetric.MType)
	if mType != metric.COUNTER && mType != metric.GAUGE {
		fmt.Println("invalid_type")
		http.Error(w, "invalid_type", http.StatusNotImplemented)
		return
	}

	m, exist := h.metricsStorage.Get(inputMetric.ID)
	if !exist {
		http.Error(w, "metric not found", http.StatusNotFound)
		return
	}

	if h.config.IsToSign() {
		err := m.SetHash(h.config.GetKey())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	jsonBody, err := json.Marshal(&m)
	if err != nil {
		http.Error(w, fmt.Sprintf("err: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	h.logger.Infof("req [%s:%s] output: [%s:%s] %s", inputMetric.ID, inputMetric.MType, m.GetName(), m.GetType(), m.ToString())
	w.Write(jsonBody)
}

// GetMetric - отдает одну метрику запрощенную методом GET как текст.
func (h *handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("http GetMetric")
	mType := metric.MType(chi.URLParam(r, "type"))
	if mType != metric.COUNTER && mType != metric.GAUGE {
		http.Error(w, "invalid_type", http.StatusNotImplemented)
		return
	}

	metricName := chi.URLParam(r, "name")
	m, exist := h.metricsStorage.Get(metricName)
	if !exist {
		http.Error(w, "metric not found", http.StatusNotFound)
		return
	}
	w.Write([]byte(m.ToString()))
}
