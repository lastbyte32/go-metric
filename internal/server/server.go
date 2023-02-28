package server

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lastbyte32/go-metric/internal/server/handlers"
	"github.com/lastbyte32/go-metric/internal/storage"
	"log"
	"net/http"
	"os"
)

type server struct {
	http *http.Server
	ctx  context.Context
}

func NewServer(config Configurator, ctx context.Context) *server {

	ms := storage.NewMemoryStorage(
		storage.WithContext(ctx),
		//storage.WithStore(config.getStoreFile(), config.getStoreInterval()),
		//storage.WithRestore(config.getStoreFile(), config.IsRestore()),

		storage.WithStore("./devops-metrics-db.json", 0),
		//storage.WithRestore("./devops-metrics-db.json", config.IsRestore()),
	)

	handler := handlers.NewHandler(ms)
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Compress(5),
		middleware.Recoverer,
	)

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
		ctx:  ctx,
	}
}

func (s *server) Run() error {

	go func() {
		<-s.ctx.Done()
		if err := s.http.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
	}()

	err := s.http.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("HTTP server stopped successfully")
		os.Exit(0)
	} else {
		return err
	}

	return nil
}
