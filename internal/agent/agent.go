package agent

import (
	"fmt"
	"time"
)

func Run(config Configurator) error {
	fmt.Println("Agent start")
	var (
		pollCount = int64(0)

		pollInterval   = config.getPollInterval()
		reportInterval = config.getPollInterval()

		counterMetrics []counter
		gaugeMetrics   []gauge
		allMetrics     []metric

		reportTimer = time.NewTicker(reportInterval)
		poolTimer   = time.NewTicker(pollInterval)
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
					return fmt.Errorf("Metric send err: %v", err)
				}
			}
		}
	}
}
