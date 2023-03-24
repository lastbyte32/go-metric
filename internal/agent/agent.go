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

func (a *agent) sender(ctx context.Context, num int) {
	a.logger.Info("start sender #", num)
	for {
		select {
		case <-ctx.Done():
			a.logger.Info("stop sender #", num)
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
		fmt.Println("sender err: ", err)
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

func (a *agent) Pool() {
	memStats := getMemStat()

	for n, v := range getCPU() {
		memStats[n] = v
	}
	for n, v := range getMemory() {
		memStats[n] = v
	}

	for n, v := range memStats {
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

func (a *agent) Run(ctx context.Context) {
	a.logger.Info("Agent start")
	reportTimer := time.NewTicker(a.config.getReportInterval())
	poolTimer := time.NewTicker(a.config.getPollInterval())

	defer func() {
		poolTimer.Stop()
		reportTimer.Stop()
	}()

	for i := 0; i < a.config.getRateLimit(); i++ {
		senderNum := i
		go a.sender(ctx, senderNum)
	}

	for {
		select {
		case <-ctx.Done():
			close(a.ReqCh)
			a.logger.Info("shutdown agent")
			return
		case <-poolTimer.C:
			a.Pool()
		case <-reportTimer.C:
			a.sendReport()
			a.sendAllReport()
		}
	}
}
