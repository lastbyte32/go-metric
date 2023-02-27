package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lastbyte32/go-metric/internal/server/handlers"
	"github.com/lastbyte32/go-metric/internal/storage"
	"net/http"
)

type server struct {
	http *http.Server
}

func NewServer(config Configurator) *server {
	ms := storage.NewMemoryStorage()
	handler := handlers.NewHandler(ms)
	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.Recoverer)
	router.Group(func(r chi.Router) {
		r.Get("/", handler.Index)
		r.Get("/value/{type}/{name}", handler.GetMetric)
		r.Post("/update/{type}/{name}/{value}", handler.UpdateMetric)
		r.Post("/update/", handler.UpdateMetricFromJSON)
		r.Post("/value/", handler.GetMetricFromJSON)
	})

	srv := &http.Server{
		Addr:    config.getAddress(),
		Handler: router,
	}
	return &server{
		http: srv,
	}
}

func (s *server) Run() error {

	err := s.http.ListenAndServe()
	if err != nil {
		return err
	}

	fmt.Println("Server start")
	return nil
}
