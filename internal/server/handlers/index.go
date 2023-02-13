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

func (h *Main) Index(w http.ResponseWriter, r *http.Request) {
	values := map[string]string{}
	body := ""
	var keys []string

	for name, value := range h.GaugeStorage.All() {
		values[name] = fmt.Sprintf("%f", value)
	}
	for name, value := range h.CountersStorage.All() {
		values[name] = fmt.Sprintf("%d", value)
	}

	if len(values) == 0 {
		body = "<li><b>No metrics</b></li>"
		responseIndex(w, body)
		return
	}

	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, val := range keys {
		body += fmt.Sprintf("<li><b>%s:</b> %s</li>", val, values[val])
	}
	responseIndex(w, body)
}
