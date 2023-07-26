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

var buildVersion = "N/A"
var buildDate = "N/A"
var buildCommit = "N/A"

func main() {
	printBuildInfo()
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
	app, err := agent.NewAgent(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	app.Run(ctx)
}

func printBuildInfo() {
	log.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n",
		buildVersion,
		buildDate,
		buildDate,
	)
}
