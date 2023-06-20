package handlers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Metrics struct {
	ID    string  `json:"id"`              // имя метрики
	MType string  `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func ExampleGetMetricFromJSON() {
	metric := Metrics{
		ID:    "test_counter",
		MType: "counter",
		Delta: 100,
	}
	bodyJSON, err := json.Marshal(metric)
	if err != nil {
		log.Fatalf("error on marshal: %v", err)
	}
	req, err := http.Post("http://example.com/value/", "application/json", bytes.NewBuffer(bodyJSON))
	defer req.Body.Close()

	if err != nil {
		log.Fatalf("error in get metric: %v", err)
	}
	log.Printf("get metric from json successful")
}
