package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	envPostgresHost     = "POSTGRES_HOST"
	envPostgresPort     = "POSTGRES_PORT"
	envPostgresUser     = "POSTGRES_USER"
	envPostgresPassword = "POSTGRES_PASSWORD"
	envPostgresDatabase = "POSTGRES_DATABASE"
)

type PostgresConfig struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
}

func NewPostgresConfig() (*PostgresConfig, error) {
	envs := []string{
		envPostgresHost,
		envPostgresPort,
		envPostgresUser,
		envPostgresPassword,
		envPostgresDatabase,
	}

	values := make(map[string]string)

	for _, env := range envs {
		value := os.Getenv(env)

		if len(value) == 0 {
			return nil, errors.New(fmt.Sprintf("%s not found", env))
		}

		values[env] = value
	}

	port, err := strconv.ParseUint(values[envPostgresPort], 10, 16)
	if err != nil || port < 0 || port > 65535 {
		return nil, errors.New(fmt.Sprintf("%s is incorrect", envPostgresPort))
	}

	return &PostgresConfig{
		Host:     values[envPostgresHost],
		Port:     uint16(port),
		User:     values[envPostgresUser],
		Password: values[envPostgresPassword],
		Database: values[envPostgresDatabase],
	}, nil
}
