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
	ms := storage.NewMemoryStorage()

	handler := handlers.NewHandler(ms)
	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.Recoverer)
	router.Group(func(r chi.Router) {
		r.Get("/", handler.Index)
		r.Get("/value/{type}/{name}", handler.GetMetric)
		r.Post("/update/{type}/{name}/{value}", handler.UpdateMetric)
	})

	srv := &http.Server{
		Addr:    config.getAddress(),
		Handler: router,
	}

	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return errors.New("server closed")
	} else if err != nil {
		return err
	}

	fmt.Println("Server listen on " + config.getAddress())
	return nil
}
