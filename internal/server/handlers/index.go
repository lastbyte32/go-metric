package handlers

import (
	"fmt"
	"net/http"
	"sort"
)

func responseIndex(w http.ResponseWriter, body string) {
	header := "<html><head><title>Metrics</title></head><body><ul>"
	footer := "</ul></body></html>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(header + body + footer))
	//return
	//w.WriteHeader(status)
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	//values := map[string]string{}
	body := ""

	metrics := h.metricsStorage.All()

	if len(metrics) == 0 {
		body = "<li><b>No metrics</b></li>"
		responseIndex(w, body)
		return
	}

	var keys []string

	for key := range metrics {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, k := range keys {
		metric := metrics[k]
		body += fmt.Sprintf("<li><b>%s:</b> %s</li>", metric.GetName(), metric.ToString())
	}
	responseIndex(w, body)
}
