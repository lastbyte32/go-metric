package agent

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
)

type IConfigurator interface {
	getAddress() string
	getReportInterval() time.Duration
	getReportTimeout() time.Duration
	getPollInterval() time.Duration
}

const (
	addressDefault        = "127.0.0.1:8080"
	reportIntervalDefault = 10 * time.Second
	pollIntervalDefault   = 2 * time.Second
	reportTimeoutDefault  = 20 * time.Second
)

type config struct {
	Address        string        `env:"ADDRESS"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	ReportTimeout  time.Duration `env:"REPORT_TIMEOUT" envDefault:"20s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL"`
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

func (c *config) flags() {
	flag.StringVar(&c.Address, "a", addressDefault, "metric server address")
	flag.DurationVar(&c.ReportInterval, "r", reportIntervalDefault, "report interval")
	flag.DurationVar(&c.PollInterval, "p", pollIntervalDefault, "poll interval")
	flag.DurationVar(&c.ReportTimeout, "t", reportTimeoutDefault, "report timeout")
	flag.Parse()
}

func NewConfig() (IConfigurator, error) {
	c := &config{}
	c.flags()
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
