package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
)

type Configurator interface {
	GetAddress() string
	GetStoreInterval() time.Duration
	GetStoreFile() string
	IsRestore() bool
	GetKey() string
	IsToSign() bool
	GetDSN() string
}

const (
	addressDefault       = "127.0.0.1:8080"
	storeIntervalDefault = 10 * time.Second
	storeFileDefault     = "/tmp/devops-metrics-db.json"
	restoreDefault       = false
	keyDefault           = ""
	DatabaseDSNDefault   = ""
)

type config struct {
	Address       string        `env:"ADDRESS"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	StoreFile     string        `env:"STORE_FILE"`
	Restore       bool          `env:"RESTORE"`
	Key           string        `env:"KEY"`
	DSN           string        `env:"DATABASE_DSN"`
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

func (c *config) GetKey() string {
	return c.Key
}

func (c *config) GetDSN() string {
	return c.DSN
}

func (c *config) IsToSign() bool {
	return c.Key != ""
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
	flag.StringVar(&c.Key, "k", keyDefault, "key for encrypt")
	flag.StringVar(&c.DSN, "d", DatabaseDSNDefault, "dsn")

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
