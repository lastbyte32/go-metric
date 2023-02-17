package server

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lastbyte32/go-metric/internal/server/handlers"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"net/http"
)

func Run(config Configurator) error {

	handler := handlers.NewHandler(storage.NewMemoryStorage())

	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.Recoverer)
	router.Group(func(r chi.Router) {
		r.Get("/", handler.Index)
		r.Get("/value/{type:gauge|counter}/{name}", handler.GetMetric)
		r.Post("/update/{type:gauge|counter}/{name}/{value}", handler.UpdateMetric)
	})

	srv := &http.Server{Addr: config.getAddress(), Handler: router}

	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server error: %v", err)
	}
	fmt.Println("Server listen on " + config.getAddress())
	return nil
}
