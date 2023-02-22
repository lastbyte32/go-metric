package agent

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lastbyte32/go-metric/internal/metric"
	"sync"
	"time"
)

type agent struct {
	reportTimer *time.Ticker
	poolTimer   *time.Ticker
	pollCount   int64
	metrics     map[string]metric.Metric
	client      *resty.Client
	config      Configurator
	sync.Mutex
}

func NewAgent(c Configurator) *agent {
	return &agent{
		reportTimer: time.NewTicker(c.getReportInterval()),
		poolTimer:   time.NewTicker(c.getReportInterval()),
		pollCount:   int64(0),
		metrics:     map[string]metric.Metric{},
		client:      resty.New().SetTimeout(c.getReportTimeout()),
		config:      c,
	}
}

func (a *agent) transmitPlainText(m metric.Metric) error {
	url := fmt.Sprintf("http://%s/update/%s/%s/%s", a.config.getAddress(), m.GetType(), m.GetName(), m.ToString())

	_, err := a.client.R().
		SetHeader("Content-Type", "text/plain").
		SetBody(m.ToString()).
		Post(url)
	if err != nil {
		return err
	}
	return nil
}

func (a *agent) sendReport() {
	fmt.Println("reportTimer")
	for _, m := range a.metrics {
		err := a.transmitPlainText(m)
		if err != nil {
			fmt.Printf("metric send err: %v", err)
		}
	}
}

func (a *agent) Pool() {
	fmt.Println("Pool...")
	memstat := getMemStat()
	a.Lock()
	defer a.Unlock()
	for n, v := range memstat {
		a.metrics[n] = metric.NewGauge(n, v)
	}
	a.pollCount++
	a.metrics["PollCount"] = metric.NewCounter("PollCount", a.pollCount)
}

func (a *agent) Run() {
	fmt.Println("Agent start")

	defer func() {
		a.poolTimer.Stop()
		a.reportTimer.Stop()
	}()

	for {
		select {
		case <-a.poolTimer.C:
			a.Pool()
		case <-a.reportTimer.C:
			a.sendReport()
		}
	}
}
