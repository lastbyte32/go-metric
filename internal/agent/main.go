package agent

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("Agent start")
	var (
		pollCount = int64(0)

		pollInterval   = 2 * time.Second
		reportInterval = 10 * time.Second

		counterMetrics []counter
		gaugeMetrics   []gauge
		reportTimer    = time.NewTicker(reportInterval)
		poolTimer      = time.NewTicker(pollInterval)
	)

	defer func() {
		poolTimer.Stop()
		reportTimer.Stop()
	}()

	for {
		select {
		case <-poolTimer.C:
			pollCount++
			poolCounter(&counterMetrics, pollCount)
			poolGauge(&gaugeMetrics)

		case <-reportTimer.C:
			sendReport(&counterMetrics, &gaugeMetrics)
		}
	}
}
