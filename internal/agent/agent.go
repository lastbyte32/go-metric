package agent

import (
	"fmt"
	"time"
)

func Run(config Configurator) {
	fmt.Println("Agent start")
	var (
		pollCount = int64(0)

		counterMetrics []counter
		gaugeMetrics   []gauge
		allMetrics     []metric

		reportTimer = time.NewTicker(config.getReportInterval())
		poolTimer   = time.NewTicker(config.getPollInterval())
	)

	defer func() {
		poolTimer.Stop()
		reportTimer.Stop()
	}()

	for {
		select {
		case <-poolTimer.C:
			pollCount++
			counterMetrics = poolCounter(pollCount)
			gaugeMetrics = poolGauge()

			allMetrics = nil

			for _, c := range counterMetrics {
				allMetrics = append(allMetrics, c)
			}
			for _, g := range gaugeMetrics {
				allMetrics = append(allMetrics, g)
			}

		case <-reportTimer.C:

			for _, m := range allMetrics {
				err := m.sendReport(config.getAddress(), config.getReportTimeout())
				if err != nil {
					fmt.Printf("metric send err: %v", err)
				}
			}
		}
	}
}
