package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lastbyte32/go-metric/internal/server"
)

func main() {
	config, err := server.NewConfig()
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

	app := server.NewServer(config)
	log.Fatal(app.Run(ctx))
}
