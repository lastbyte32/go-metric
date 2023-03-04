package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, ctxCancel := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	defer ctxCancel()

	app := server.NewServer(cfg)
	log.Fatal(app.Run(ctx))
}
