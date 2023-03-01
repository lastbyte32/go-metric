package server

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lastbyte32/go-metric/internal/server/handlers"
	"github.com/lastbyte32/go-metric/internal/storage"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

type server struct {
	http   *http.Server
	ctx    context.Context
	logger *zap.SugaredLogger
}

func NewServer(config Configurator, ctx context.Context) *server {

	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error on create logger: %v", err)
	}
	logger := l.Sugar()
	defer logger.Sync()

	ms := storage.NewMemoryStorage(
		storage.WithContext(ctx),
		storage.WithStore(config.getStoreFile(), config.getStoreInterval()),
		storage.WithRestore(config.getStoreFile(), config.IsRestore()),
	)

	handler := handlers.NewHandler(ms, logger)
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
		http:   srv,
		ctx:    ctx,
		logger: logger,
	}
}

func (s *server) Run() error {
	s.logger.Info("http server run")

	go func() {
		<-s.ctx.Done()
		if err := s.http.Shutdown(context.Background()); err != nil {
			s.logger.Errorf("HTTP server shutdown error: %v", err)
		}
	}()
	err := s.http.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		s.logger.Info("HTTP server stopped successfully")
		os.Exit(0)
	} else {
		return err
	}

	return nil
}
