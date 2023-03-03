package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
)

type Configurator interface {
	IStorage
	IServer
	IAgent
}

type IStorage interface {
	IsRestore() bool
	GetStoreFile() string
	GetStoreInterval() time.Duration
}

type IAgent interface {
	GetAddress() string
	GetReportInterval() time.Duration
	GetReportTimeout() time.Duration
	GetPollInterval() time.Duration
}

type IServer interface {
	GetAddress() string
}

const (
	addressDefault        = "127.0.0.1:8080"
	storeIntervalDefault  = 10 * time.Second
	storeFileDefault      = "/tmp/devops-metrics-db.json"
	restoreDefault        = false
	reportIntervalDefault = 10 * time.Second
	pollIntervalDefault   = 2 * time.Second
	reportTimeoutDefault  = 20 * time.Second
)

type config struct {
	Address        string        `env:"ADDRESS"`
	StoreInterval  time.Duration `env:"STORE_INTERVAL"`
	StoreFile      string        `env:"STORE_FILE"`
	Restore        bool          `env:"RESTORE"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	ReportTimeout  time.Duration `env:"REPORT_TIMEOUT" envDefault:"20s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL"`
}

func (c *config) GetStoreFile() string {
	return c.StoreFile
}

func (c *config) IsRestore() bool {
	return c.Restore
}

func (c *config) GetAddress() string {
	return c.Address
}

func (c *config) GetStoreInterval() time.Duration {
	return c.StoreInterval
}

func (c *config) GetReportInterval() time.Duration {
	return c.ReportInterval
}

func (c *config) GetReportTimeout() time.Duration {
	return c.ReportTimeout
}

func (c *config) GetPollInterval() time.Duration {
	return c.PollInterval
}

func (c *config) env() error {
	if err := env.Parse(c); err != nil {
		return err
	}
	return nil
}

func (c *config) flags() {
	flag.StringVar(&c.Address, "a", addressDefault, "server binding host:port")
	flag.StringVar(&c.StoreFile, "f", storeFileDefault, "store metrics in file")
	flag.BoolVar(&c.Restore, "r", restoreDefault, "restoreDefault metrics")
	flag.DurationVar(&c.StoreInterval, "i", storeIntervalDefault, "store metrics on interval")
	flag.DurationVar(&c.ReportInterval, "r", reportIntervalDefault, "report interval")
	flag.DurationVar(&c.PollInterval, "p", pollIntervalDefault, "poll interval")
	flag.DurationVar(&c.ReportTimeout, "t", reportTimeoutDefault, "report timeout")
	flag.Parse()
}

func NewConfig() (Configurator, error) {
	c := &config{}
	c.flags()
	err := c.env()
	if err != nil {
		return nil, err
	}
	fmt.Printf("*Configuration used*\n\t- Address: %s\n\t- StoreInterval: %.0fs\n",
		c.Address,
		c.StoreInterval.Seconds(),
	)
	return c, nil
}
