package agent

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lastbyte32/go-metric/internal/metric"
	"github.com/lastbyte32/go-metric/internal/storage"
	"sync"
	"time"
)

type agent struct {
	ms     storage.Storage
	client *resty.Client
	config Configurator
	sync.Mutex
}

func NewAgent(config Configurator) *agent {

	client := resty.New().
		SetTimeout(config.getReportTimeout())
	memory := storage.NewMemoryStorage()
	return &agent{
		client: client,
		config: config,
		ms:     memory,
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
	fmt.Println("sendReport")
	for _, m := range a.ms.All() {
		err := a.transmitPlainText(m)
		if err != nil {
			fmt.Printf("err send [%s]: %v\n", m.GetName(), err)
		}
	}
}

func (a *agent) Pool() {
	fmt.Println("Pool...")
	memstat := getMemStat()
	a.Lock()
	defer a.Unlock()
	for n, v := range memstat {
		value := fmt.Sprintf("%.3f", v)
		err := a.ms.Update(n, value, metric.GAUGE)
		if err != nil {
			fmt.Printf("err update %s", n)
		}
	}

	err := a.ms.Update("PollCount", "1", metric.COUNTER)
	if err != nil {
		fmt.Printf("err update %s", "PollCount")
	}
}

func (a *agent) Run() {
	fmt.Println("Agent start")
	reportTimer := time.NewTicker(a.config.getReportInterval())
	poolTimer := time.NewTicker(a.config.getPollInterval())

	defer func() {
		poolTimer.Stop()
		reportTimer.Stop()
	}()

	for {
		select {
		case <-poolTimer.C:
			a.Pool()
		case <-reportTimer.C:
			a.sendReport()
		}
	}
}
