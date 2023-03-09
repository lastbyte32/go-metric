package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/lastbyte32/go-metric/internal/metric"
	"github.com/lastbyte32/go-metric/internal/utils"
)

func (h *handler) UpdateMetricFromJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m metric.Metrics
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mtype := metric.MType(m.MType)
	if mtype != metric.COUNTER && mtype != metric.GAUGE {
		http.Error(w, "invalid_type", http.StatusNotImplemented)
		return
	}

	if h.config.IsToSign() && m.Hash != "" {
		isVerify, err := m.VerifyHash(h.config.GetKey())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !isVerify {
			http.Error(w, "hash not equal", http.StatusBadRequest)
			return
		}
	}

	var value = ""
	switch mtype {
	case metric.COUNTER:
		value = fmt.Sprintf("%d", *m.Delta)
	case metric.GAUGE:
		value = utils.FloatToString(*m.Value)
	default:
		fmt.Println("invalid_type")
		http.Error(w, "invalid_type", http.StatusNotImplemented)
		return
	}

	err := h.metricsStorage.Update(m.ID, value, mtype)
	if err != nil {
		http.Error(w, fmt.Sprintf("err: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// тут пропускаем ошибку потому что выше проверили удачный ли апдейт
	updatedMetric, _ := h.metricsStorage.Get(m.ID)
	if h.config.IsToSign() {
		err := updatedMetric.SetHash(h.config.GetKey())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	jsonBody, err := json.Marshal(updatedMetric)
	if err != nil {
		http.Error(w, fmt.Sprintf("err: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Write(jsonBody)
}

func (h *handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateHandle")
	metricType := metric.MType(chi.URLParam(r, "type"))
	metricName := chi.URLParam(r, "name")

	if metricType != metric.COUNTER && metricType != metric.GAUGE {
		fmt.Println("invalid_type")
		http.Error(w, "invalid_type", http.StatusNotImplemented)
		return
	}

	value := chi.URLParam(r, "value")
	err := h.metricsStorage.Update(metricName, value, metricType)
	if err != nil {
		http.Error(w, fmt.Sprintf("err: %s", err.Error()), http.StatusBadRequest)
		return
	}
}
