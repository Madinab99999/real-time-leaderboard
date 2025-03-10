package config

import (
	"context"
	"flag"
	"os"
	"strconv"
)

type APIConfig struct {
	Rest *APIRestConfig
}

func newApiConfig(ctx context.Context) *APIConfig {
	c := &APIConfig{

		Rest: newApiRestConfig(ctx),
	}

	return c
}

type APIRestConfig struct {
	Host string
	Port int
}

func newApiRestConfig(_ context.Context) *APIRestConfig {
	port, _ := strconv.Atoi(os.Getenv("API_REST_PORT"))

	c := &APIRestConfig{
		Host: os.Getenv("API_REST_HOST"),
		Port: port,
	}

	flag.StringVar(&c.Host, "api-rest-host", c.Host, "api rest host [API_REST_HOST]")
	flag.IntVar(&c.Port, "api-rest-port", c.Port, "api rest port [API_REST_PORT]")

	return c
}
