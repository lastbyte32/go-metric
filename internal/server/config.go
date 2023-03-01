package server

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
)

type Configurator interface {
	getAddress() string
	getStoreInterval() time.Duration
	getStoreFile() string
	IsRestore() bool
}

const (
	address       = "127.0.0.1:8080"
	storeInterval = 300 * time.Second
	storeFile     = "/tmp/devops-metrics-db.json"
	restore       = true
)

type config struct {
	Address       string        `env:"ADDRESS"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	StoreFile     string        `env:"STORE_FILE"`
	Restore       bool          `env:"RESTORE"`
}

func (c *config) getStoreFile() string {
	return c.StoreFile
}

func (c *config) IsRestore() bool {
	return c.Restore
}

func (c *config) getAddress() string {
	return c.Address
}

func (c *config) getStoreInterval() time.Duration {
	return c.StoreInterval
}

func (c *config) env() error {
	if err := env.Parse(c); err != nil {
		return err
	}
	return nil
}

func (c *config) flags() {
	flag.StringVar(&c.Address, "a", address, "server binding host:port")
	flag.StringVar(&c.StoreFile, "f", storeFile, "store metrics in file")
	flag.BoolVar(&c.Restore, "r", restore, "restore metrics")
	flag.DurationVar(&c.StoreInterval, "i", storeInterval, "store metrics on interval")
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
