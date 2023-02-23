package server

import (
	"fmt"
)

// Configurator Todo: подумать на тем что бы в дальнейшем сделать
type Configurator interface {
	getAddress() string
}

type config struct {
	address string
}

func (c *config) getAddress() string {
	return c.address
}

func (c *config) defaultConfigParam() {
	const address = ":8080"
	c.address = address
}

func NewConfig() (error, *config) {
	//Todo: Реализовать загрузку "конфига" из файла/флагов/окружения
	c := &config{}
	c.defaultConfigParam()
	fmt.Printf("*Configuration used*\n\t- Address: %s\n", c.address)
	return nil, c
}
