package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/server/handlers"
	customMiddleware "github.com/lastbyte32/go-metric/internal/server/middleware"
	"github.com/lastbyte32/go-metric/internal/storage"
)

type server struct {
	http   *http.Server
	store  storage.IStorage
	logger *zap.SugaredLogger
}

func NewServer(config config.Configurator) *server {

	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error on create logger: %v", err)
	}
	logger := l.Sugar()
	defer logger.Sync()

	store := storage.New(config, logger)

	handler := handlers.NewHandler(store, logger, config)
	router := chi.NewRouter()
	router.Use(
		customMiddleware.Compressor,
		middleware.Logger,
		//	middleware.Recoverer,
	)

	router.Group(func(r chi.Router) {
		r.Get("/", handler.Index)
		r.Get("/value/{type}/{name}", handler.GetMetric)
		r.Post("/update/{type}/{name}/{value}", handler.UpdateMetric)
		r.Post("/update/", handler.UpdateMetricFromJSON)
		r.Post("/updates/", handler.UpdatesMetricFromJSON)
		r.Post("/value/", handler.GetMetricFromJSON)
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if config.GetDSN() == "" {
			http.Error(w, "no db", http.StatusNotImplemented)
			return
		}

		if err := storage.Ping(config.GetDSN()); err != nil {
			http.Error(w, "db err", http.StatusInternalServerError)
			return
		}
	})

	httpServer := &http.Server{
		Addr:    config.GetAddress(),
		Handler: router,
	}
	return &server{
		http:   httpServer,
		logger: logger,
		store:  store,
	}
}

func (s *server) shutdown() {
	if err := s.store.Close(); err != nil {
		s.logger.Errorf("store shutdown error: %v", err)
	}

	if err := s.http.Shutdown(context.Background()); err != nil {
		s.logger.Errorf("http server shutdown error: %v", err)
	}
	s.logger.Info("shutdown completed")
}

func (s *server) Run(ctx context.Context) error {
	go func() {
		s.logger.Info("start shutdown watcher")
		<-ctx.Done()
		s.logger.Info("Received signal, stopping application")
		s.shutdown()
	}()
	s.logger.Info("http server run")
	err := s.http.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		s.logger.Info("HTTP server stopped successfully")
		os.Exit(0)
	} else {
		return err
	}

	return nil
}
