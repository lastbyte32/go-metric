package agent

import (
	"fmt"
	"time"
)

const (
	address        = "127.0.0.1:8080"
	reportInterval = 10
	pollInterval   = 2
	reportTimeout  = 20
)

type Configurator interface {
	getAddress() string
	getReportInterval() time.Duration
	getReportTimeout() time.Duration
	getPollInterval() time.Duration
}

type config struct {
	address        string
	reportInterval time.Duration
	reportTimeout  time.Duration
	pollInterval   time.Duration
}

func (c *config) getAddress() string {
	return c.address
}

func (c *config) getReportInterval() time.Duration {
	return c.reportInterval
}

func (c *config) getReportTimeout() time.Duration {
	return c.reportTimeout
}

func (c *config) getPollInterval() time.Duration {
	return c.pollInterval
}

func (c *config) defaultConfigParam() {
	c.address = address
	c.reportInterval = time.Second * reportInterval
	c.reportTimeout = time.Second * reportTimeout
	c.pollInterval = time.Second * pollInterval
}

func NewConfig() (Configurator, error) {
	//Todo: Реализовать загрузку "конфига" из файла/флагов/окружения
	c := &config{}
	c.defaultConfigParam()
	fmt.Printf("*Configuration used*\n\t- Server: %s\n\t- ReportInterval: %.0fs\n\t- PollInterval: %.0fs\n", c.address, c.reportInterval.Seconds(), c.pollInterval.Seconds())
	return c, nil
}
