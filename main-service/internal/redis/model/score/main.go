package score

import (
	"log/slog"

	"main_service/internal/config"

	"github.com/redis/go-redis/v9"
)

type Score struct {
	conf   *config.RedisConfig
	logger *slog.Logger
	db     *redis.Client
}

func New(conf *config.RedisConfig, logger *slog.Logger, db *redis.Client) *Score {
	return &Score{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
