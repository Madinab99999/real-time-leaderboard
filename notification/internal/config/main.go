package config

import (
	"context"
	"flag"
	"os"

	"notification/internal/constant"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV   constant.Environment
	REDIS *RedisConfig
}

func New(ctx context.Context) *Config {
	var env constant.Environment
	e := os.Getenv("ENV")
	if e == "" {
		env = constant.EnvironmentLocal
	} else {
		env = constant.Environment(e)
	}

	load(env)

	conf := &Config{
		ENV:   env,
		REDIS: newRedisConfig(ctx),
	}

	flag.Parse()

	return conf
}

func load(env constant.Environment) error {
	switch env {
	case constant.EnvironmentProd:
		return godotenv.Load(".env", ".env.local", ".env.prod")
	case constant.EnvironmentDev:
		return godotenv.Load(".env", ".env.local", ".env.dev")
	case constant.EnvironmentTest:
		return godotenv.Load(".env", ".env.local", ".env.test")
	case constant.EnvironmentLocal:
		return godotenv.Load(".env", ".env.local")
	default:
		return nil
	}
}
