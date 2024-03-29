package handlers

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

func Example_handler_GetMetricFromJSON() {
	metric := Metrics{
		ID:    "test_counter",
		MType: "counter",
	}
	bodyJSON, err := json.Marshal(metric)
	if err != nil {
		log.Fatalf("error on marshal: %v", err)
	}
	req, err := http.Post("http://example.com/value/", "application/json", bytes.NewBuffer(bodyJSON))
	if err != nil {
		log.Fatalf("error in get metric: %v", err)
	}

	err = req.Body.Close()
	if err != nil {
		log.Fatalf("error on close: %v", err)
	}
	log.Printf("get metric from json successful")
}
