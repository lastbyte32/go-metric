package agent

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v7"
)

type IConfigurator interface {
	getAddress() string
	getReportInterval() time.Duration
	getReportTimeout() time.Duration
	getPollInterval() time.Duration
	getKey() string
	isToSign() bool
	getRateLimit() int
	GetCryptoKeyPath() string
}

const (
	addressDefault        = "127.0.0.1:8080"
	reportIntervalDefault = 10 * time.Second
	pollIntervalDefault   = 2 * time.Second
	reportTimeoutDefault  = 20 * time.Second
	keyDefault            = ""
	RateLimitDefault      = 10
	cryptoKeyPathDefault  = ""
	configPathDefault     = ""
)

type config struct {
	Address        string        `env:"ADDRESS" json:"address"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" json:"report_interval"`
	ReportTimeout  time.Duration `env:"REPORT_TIMEOUT" envDefault:"20s" json:"report_timeout,omitempty"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" json:"poll_interval"`
	Key            string        `env:"KEY" json:"key"`
	RateLimit      int           `env:"RATE_LIMIT" json:"rate_limit,omitempty"`
	CryptoKeyPath  string        `env:"CRYPTO_KEY" json:"crypto_key"`
	ConfigPath     string        `env:"CONFIG"`
}

// GetCryptoKeyPath - метод возвращает путь до файла с публичным ключом.
func (c *config) GetCryptoKeyPath() string {
	return c.CryptoKeyPath
}

func (c *config) getRateLimit() int {
	return c.RateLimit
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

func (c *config) getKey() string {
	return c.Key
}

func (c *config) isToSign() bool {
	return c.Key != ""
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
	flag.StringVar(&c.Key, "k", keyDefault, "key for encrypt")
	flag.IntVar(&c.RateLimit, "l", RateLimitDefault, "RateLimit")
	flag.StringVar(&c.CryptoKeyPath, "crypto-key", cryptoKeyPathDefault, "path to public key")
	flag.StringVar(&c.ConfigPath, "c", configPathDefault, "path to json config file")
	flag.Parse()
}

func NewConfig() (IConfigurator, error) {
	c := &config{}
	c.flags()
	if err := c.env(); err != nil {
		return nil, err
	}

	if c.ConfigPath != "" {
		file, err := os.ReadFile(c.ConfigPath)
		if err != nil {
			return nil, err
		}
		var jsonCfg config
		if err := json.Unmarshal(file, &jsonCfg); err != nil {
			return nil, err
		}
		c = &jsonCfg
	}

	fmt.Printf("*Configuration used*\n\t- Server: %s\n\t- ReportInterval: %.0fs\n\t- PollInterval: %.0fs\n",
		c.Address,
		c.ReportInterval.Seconds(),
		c.PollInterval.Seconds(),
	)
	return c, nil
}
