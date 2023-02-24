package main

import (
	"github.com/lastbyte32/go-metric/internal/server"
	"log"
)

func main() {
	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	app := server.NewServer(config)
	err = app.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
