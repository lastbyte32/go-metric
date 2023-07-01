package main

import (
	"context"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/lastbyte32/go-metric/internal/agent"
	"github.com/lastbyte32/go-metric/pkg/utils/profile"
)

func main() {
	profile.ProfileIfEnabled()
	config, err := agent.NewConfig()
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
	app := agent.NewAgent(config)
	app.Run(ctx)
}
