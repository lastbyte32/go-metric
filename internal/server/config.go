package server

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"time"
)

type Configurator interface {
	getAddress() string
	getStoreInterval() time.Duration
	getStoreFile() string
	IsRestore() bool
}

type config struct {
	Address       string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
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

func NewConfig() (Configurator, error) {
	c := &config{}
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
