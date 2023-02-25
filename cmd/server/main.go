package main

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/server"
	"log"
)

func main() {
	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = server.Run(config)
	if err != nil {
		fmt.Printf("metric server err: %v", err)
	}
}
