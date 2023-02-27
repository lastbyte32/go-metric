package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/metric"
	"net/http"
)

func (h *handler) UpdateMetricFromJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsonBody")

	var m metric.Metrics
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mtype := metric.MType(m.MType)
	if mtype != metric.COUNTER && mtype != metric.GAUGE {
		fmt.Println("invalid_type")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("err invalid_type"))
		return
	}
	var value = ""
	switch mtype {
	case metric.COUNTER:
		inputValue := *m.Delta
		value = fmt.Sprintf("%d", inputValue)
		//fmt.Printf("Name: %s\nType: %s\nValue: %s\n", m.ID, m.MType, value)
	case metric.GAUGE:
		inputValue := *m.Value
		value = fmt.Sprintf("%f", inputValue)
		//fmt.Printf("Name: %s\nType: %s\nValue: %s\n", m.ID, m.MType, value)
	default:
		fmt.Println("invalid_type")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("err invalid_type"))
		return
	}

	err := h.metricsStorage.Update(m.ID, value, mtype)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// тут пропускаем ошибку потому что выше проверили удачный ли апдейт
	updatedMetric, _ := h.metricsStorage.Get(m.ID)
	jsonBody, err := updatedMetric.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
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
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("err invalid_type"))
		return
	}

	value := chi.URLParam(r, "value")
	err := h.metricsStorage.Update(metricName, value, metricType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}
