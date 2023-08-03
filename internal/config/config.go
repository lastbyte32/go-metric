package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v7"
)

// Configurator TODO: перенести объявление интерфейса по месту использования
type Configurator interface {
	GetAddress() string
	GetStoreInterval() time.Duration
	GetStoreFile() string
	IsRestore() bool
	GetKey() string
	IsToSign() bool
	GetDSN() string
	GetCryptoKeyPath() string
	GetTrustedSubnet() string
}

const (
	addressDefault       = "127.0.0.1:8080"
	storeIntervalDefault = 300 * time.Second
	storeFileDefault     = "/tmp/devops-metrics-db.json"
	restoreDefault       = false
	keyDefault           = ""
	databaseDSNDefault   = ""
	cryptoKeyPathDefault = ""
	configPathDefault    = ""
	trustedSubnetDefault = ""
)

type config struct {
	Address       string        `env:"ADDRESS" json:"address"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" json:"store_interval"`
	StoreFile     string        `env:"STORE_FILE" json:"store_file"`
	Restore       bool          `env:"RESTORE" json:"restore"`
	Key           string        `env:"KEY"`
	DSN           string        `env:"DATABASE_DSN" json:"database_dsn"`
	CryptoKeyPath string        `env:"CRYPTO_KEY" json:"crypto_key"`
	ConfigPath    string        `env:"CONFIG"`
	TrustedSubnet string        `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
}

// GetTrustedSubnet - метод возвращает доверенную подсеть.
func (c *config) GetTrustedSubnet() string {
	return c.TrustedSubnet
}

// GetCryptoKeyPath - метод возвращает путь до файла с публичным ключом.
func (c *config) GetCryptoKeyPath() string {
	return c.CryptoKeyPath
}

// GetStoreFile - метод возвращает путь до файла хранения метрик.
func (c *config) GetStoreFile() string {
	return c.StoreFile
}

// IsRestore - нужно ли восстанавливать метрики из файла, true - да, по-умолчанию метрики не восстанавливаем.
func (c *config) IsRestore() bool {
	return c.Restore
}

// GetAddress - адрес биндинга http сервер.
func (c *config) GetAddress() string {
	return c.Address
}

// GetStoreInterval - интервал сохранения метрик на в файл.
func (c *config) GetStoreInterval() time.Duration {
	return c.StoreInterval
}

// GetKey - ключ для подписи входящих данных.
func (c *config) GetKey() string {
	return c.Key
}

// GetDSN - параметры подключения к БД.
func (c *config) GetDSN() string {
	return c.DSN
}

// IsToSign - нужно ли проверять подпись.
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
	flag.StringVar(&c.DSN, "d", databaseDSNDefault, "dsn")
	flag.StringVar(&c.CryptoKeyPath, "crypto-key", cryptoKeyPathDefault, "path to private key")
	flag.StringVar(&c.ConfigPath, "c", configPathDefault, "path to json config file")
	flag.StringVar(&c.ConfigPath, "t", trustedSubnetDefault, "allowed subnet")
	flag.Parse()
}

// NewConfig - конструктор, который инициализирует конфигурацию.
func NewConfig() (Configurator, error) {
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
	fmt.Printf("*Server configuration used*\n\t- Address: %s\n\t- StoreInterval: %.0fs\n\t-StoreFile: %s\n\t-Restore: %v\n",
		c.Address,
		c.StoreInterval.Seconds(),
		c.StoreFile,
		c.Restore,
	)
	return c, nil
}
