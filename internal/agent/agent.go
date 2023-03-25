package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/metric"
	"github.com/lastbyte32/go-metric/internal/storage"
)

type Request struct {
	Body string
	URL  string
}

type agent struct {
	ms     storage.IStorage
	client *resty.Client
	config IConfigurator
	logger *zap.SugaredLogger
	ReqCh  chan *Request
}

func NewAgent(config IConfigurator) *agent {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error on create logger: %v", err)
	}
	logger := l.Sugar()
	defer logger.Sync()

	client := resty.New().
		SetTimeout(config.getReportTimeout())
	memory := storage.NewMemoryStorage(logger)
	return &agent{
		client: client,
		config: config,
		ms:     memory,
		logger: logger,
		ReqCh:  make(chan *Request),
	}
}

func (a *agent) addRequest(url, body string) {
	a.ReqCh <- &Request{
		Body: body,
		URL:  url,
	}
}

func (a *agent) transmitWorker(ctx context.Context, num int) {
	a.logger.Info("start transmitWorker #", num)
	for {
		select {
		case <-ctx.Done():
			a.logger.Info("stop transmitWorker #", num)
			return
		default:
			for job := range a.ReqCh {
				a.transmit(job)
			}
		}
	}
}

func (a *agent) transmit(job *Request) {
	_, err := a.client.R().
		SetHeader("Content-Type", "text/plain").
		SetBody(job.Body).
		Post(job.URL)
	if err != nil {
		fmt.Println("transmitWorker err: ", err)
		return
	}
}

func (a *agent) makePlainTextRequest(m metric.IMetric) error {
	url := fmt.Sprintf("http://%s/update/%s/%s/%s", a.config.getAddress(), m.GetType(), m.GetName(), m.ToString())
	a.addRequest(url, "")
	return nil
}

func (a *agent) makeJSONRequest(m metric.IMetric) error {
	url := fmt.Sprintf("http://%s/update/", a.config.getAddress())

	if a.config.isToSign() {
		err := m.SetHash(a.config.getKey())
		if err != nil {
			return err
		}
	}

	body, err := json.Marshal(&m)
	if err != nil {
		log.Printf("Error in JSON marshal. Err: %s", err)
		return err
	}
	a.addRequest(url, string(body))
	return nil
}

func (a *agent) sendAllReport() {
	url := fmt.Sprintf("http://%s/updates/", a.config.getAddress())
	all := a.ms.All()

	if a.config.isToSign() {
		for _, m := range all {
			m.SetHash(a.config.getKey())
		}
	}

	body, err := json.Marshal(&all)
	if err != nil {
		log.Printf("Error in JSON marshal. Err: %s", err)
		return
	}

	a.addRequest(url, string(body))
}

func (a *agent) sendReport() {
	fmt.Println("sendReportALL")
	for _, m := range a.ms.All() {
		err := a.makeJSONRequest(m)
		if err != nil {
			fmt.Printf("err send [%s]: %v\n", m.GetName(), err)
		}
	}
}

func (a *agent) statWorker(ctx context.Context, getStat func() map[string]float64) {
	a.logger.Info("statWorker start")
	poolTimer := time.NewTicker(a.config.getPollInterval())
	defer poolTimer.Stop()
	for {
		select {
		case <-ctx.Done():
			a.logger.Info("shutdown statWorker")
			return
		case <-poolTimer.C:
			for n, v := range getStat() {
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
	}
}
func (a *agent) Run(ctx context.Context) {
	a.logger.Info("Agent start")
	reportTimer := time.NewTicker(a.config.getReportInterval())
	defer reportTimer.Stop()

	go a.statWorker(ctx, getMemory)
	go a.statWorker(ctx, getRunTimeStat)
	go a.statWorker(ctx, getCPU)

	for i := 0; i < a.config.getRateLimit(); i++ {
		workerNum := i
		go a.transmitWorker(ctx, workerNum)
	}

	for {
		select {
		case <-ctx.Done():
			close(a.ReqCh)
			a.logger.Info("shutdown agent")
			return
		case <-reportTimer.C:
			a.sendReport()
			a.sendAllReport()
		}
	}
}
