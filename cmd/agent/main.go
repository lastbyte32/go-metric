package main

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/agent"
)

func main() {
	err := agent.Run(agent.NewConfig())
	if err != nil {
		// Из ошибок пока может возникнуть только невозможность отправки метрик на сервер
		// В задании еще не было информации что нужно делать при таком кейсе
		fmt.Printf("Agent err: %v", err)
	}
}
