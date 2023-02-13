package server

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/server/handlers"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"log"

	"net/http"
)

func Run() {
	fmt.Println("Server start")

	gaugeStorage := storage.NewGaugeMemStorage()
	countersStorage := storage.NewCounterMemStorage()

	http.HandleFunc("/update/", handlers.UpdateHandle(gaugeStorage, countersStorage))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
