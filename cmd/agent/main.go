package main

import (
	"github.com/lastbyte32/go-metric/internal/agent"
	"log"
)

func main() {
	config, err := agent.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	app := agent.NewAgent(config)
	app.Run()
}
