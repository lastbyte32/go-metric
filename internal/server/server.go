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

	"github.com/lastbyte32/go-metric/internal/server/handlers"
	customMiddleware "github.com/lastbyte32/go-metric/internal/server/middleware"
	"github.com/lastbyte32/go-metric/internal/storage"
)

type server struct {
	http   *http.Server
	store  storage.IStorage
	logger *zap.SugaredLogger
}

func NewServer(config Configurator) *server {

	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error on create logger: %v", err)
	}
	logger := l.Sugar()
	defer logger.Sync()

	var store storage.IStorage
	if config.getStoreFile() != "" {
		store = storage.NewFileStorage(
			logger,
			config.getStoreFile(),
			config.getStoreInterval(),
			config.IsRestore(),
		)
	} else {
		store = storage.NewMemoryStorage(logger)
	}

	handler := handlers.NewHandler(store, logger)
	router := chi.NewRouter()
	router.Use(
		customMiddleware.Compressor,
		middleware.Logger,
		middleware.Recoverer,
	)

	router.Group(func(r chi.Router) {
		r.Get("/", handler.Index)
		r.Get("/value/{type}/{name}", handler.GetMetric)
		r.Post("/update/{type}/{name}/{value}", handler.UpdateMetric)
		r.Post("/update/", handler.UpdateMetricFromJSON)
		r.Post("/value/", handler.GetMetricFromJSON)
	})

	httpServer := &http.Server{
		Addr:    config.getAddress(),
		Handler: router,
	}
	return &server{
		http:   httpServer,
		logger: logger,
		store:  store,
	}
}

func (s *server) shutdownHTTP() {
	s.logger.Info("shutdown http server")
	err := s.http.Shutdown(context.Background())
	if err != nil {
		s.logger.Errorf("HTTP server shutdown error: %v", err)
	}
}

func (s *server) Run(ctx context.Context) error {
	s.logger.Info("http server run")

	go func() {
		<-ctx.Done()
		s.logger.Info("Received signal, stopping application")
		s.shutdownHTTP()
	}()

	errStore := s.store.Init(ctx)
	if errStore != nil {
		return errStore
	}

	errHTTP := s.http.ListenAndServe()
	if errors.Is(errHTTP, http.ErrServerClosed) {
		s.logger.Info("HTTP server stopped successfully")
		os.Exit(0)
	} else {
		return errHTTP
	}

	return nil
}
