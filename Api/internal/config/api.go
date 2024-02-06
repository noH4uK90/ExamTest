package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	envApiHost = "API_HOST"
	envApiPort = "API_PORT"
)

type ApiConfiguration struct {
	Host string
	Port string
}

func NewApiConfiguration() (*ApiConfiguration, error) {
	envs := []string{
		envApiHost,
		envApiPort,
	}

	values := make(map[string]string)

	for _, env := range envs {
		value := os.Getenv(env)

		if len(value) == 0 {
			return nil, errors.New(fmt.Sprintf("%s not found", env))
		}

		values[env] = value
	}

	return &ApiConfiguration{
		Host: values[envApiHost],
		Port: values[envApiPort],
	}, nil
}

func (cfg *ApiConfiguration) Address() string { return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port) }
