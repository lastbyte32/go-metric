package main

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/server"
)

func main() {
	err := server.Run(server.NewConfig())
	if err != nil {
		fmt.Printf("metric server err: %v", err)

	}
}
