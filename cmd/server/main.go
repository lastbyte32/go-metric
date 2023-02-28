package main

import (
	"context"
	"errors"
	"github.com/lastbyte32/go-metric/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, ctxCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer ctxCancel()

	app := server.NewServer(config, ctx)
	err = app.Run()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server stopped successfully")
		return
	}
	if err != nil {
		log.Fatal(err.Error())
	}
}
