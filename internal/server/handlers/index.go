package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func responseIndex(w http.ResponseWriter, body string) {
	var response strings.Builder
	response.WriteString("<html><head><title>Metrics</title></head><body><ul>")
	response.WriteString(body)
	response.WriteString("</ul></body></html>")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(response.String()))
}

// Index - html страница с метриками.
func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	var body strings.Builder
	metrics := h.metricsStorage.All()
	if len(metrics) == 0 {
		body.WriteString("<li><b>No metrics</b></li>")
		responseIndex(w, body.String())
		return
	}
	var keys []string
	for key := range metrics {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		metric := metrics[k]
		body.WriteString(fmt.Sprintf("<li><b>%s:</b> %s</li>", metric.GetName(), metric.ToString()))
	}
	responseIndex(w, body.String())
}
