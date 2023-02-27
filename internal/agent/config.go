package agent

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"time"
)

type IConfigurator interface {
	getAddress() string
	getReportInterval() time.Duration
	getReportTimeout() time.Duration
	getPollInterval() time.Duration
}

// что бы воспользоваться библиотекой github.com/caarlos0/env/v7, пришлось сделать своиства экспортируемыми
type config struct {
	Address        string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	ReportTimeout  time.Duration `env:"REPORT_TIMEOUT" envDefault:"20s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
}

func (c *config) getAddress() string {
	return c.Address
}

func (c *config) getReportInterval() time.Duration {
	return c.ReportInterval
}

func (c *config) getReportTimeout() time.Duration {
	return c.ReportTimeout
}

func (c *config) getPollInterval() time.Duration {
	return c.PollInterval
}

func (c *config) env() error {
	if err := env.Parse(c); err != nil {
		return err
	}
	return nil
}

func NewConfig() (IConfigurator, error) {
	c := &config{}
	err := c.env()
	if err != nil {
		return nil, err
	}
	fmt.Printf("*Configuration used*\n\t- Server: %s\n\t- ReportInterval: %.0fs\n\t- PollInterval: %.0fs\n",
		c.Address,
		c.ReportInterval.Seconds(),
		c.PollInterval.Seconds(),
	)
	return c, nil
}
