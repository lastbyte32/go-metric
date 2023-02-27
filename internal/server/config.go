package server

import (
	"fmt"
	"github.com/caarlos0/env/v7"
)

type Configurator interface {
	getAddress() string
}

type config struct {
	Address string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}

func (c *config) getAddress() string {
	return c.Address
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
	fmt.Printf("*Configuration used*\n\t- Address: %s\n", c.Address)
	return c, nil
}
