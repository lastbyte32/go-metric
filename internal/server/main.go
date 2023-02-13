package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lastbyte32/go-metric/internal/server/handlers"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"log"

	"net/http"
)

func Run() {

	const serverPort = "8080"
	gaugeStorage := storage.NewGaugeMemStorage()
	countersStorage := storage.NewCounterMemStorage()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	handler := &handlers.Main{
		GaugeStorage:    gaugeStorage,
		CountersStorage: countersStorage,
	}

	router.Route("/", func(r chi.Router) {
		//r.Get("/favicon.ico", handler.Favicon)
		r.Get("/", handler.Index)
	})

	router.Route("/update", func(r chi.Router) {

		r.Post("/{type}/{name}/{value}", handler.Update)
		r.Post("/gauge/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotFound) })
		r.Post("/counter/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotFound) })

		r.Post("/*", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotImplemented) })
	})

	router.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", handler.GetOneMetric)
	})

	//http.HandleFunc("/update/", handlers.UpdateHandle(gaugeStorage, countersStorage))

	fmt.Println("Server listen on " + serverPort)

	log.Fatal(http.ListenAndServe(":"+serverPort, router))
}
