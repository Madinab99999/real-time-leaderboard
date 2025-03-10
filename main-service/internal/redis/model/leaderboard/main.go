package leaderboard

import (
	"log/slog"
	"main_service/internal/config"

	"github.com/redis/go-redis/v9"
)

type Leaderboard struct {
	conf   *config.RedisConfig
	logger *slog.Logger
	db     *redis.Client
}

func New(conf *config.RedisConfig, logger *slog.Logger, db *redis.Client) *Leaderboard {
	return &Leaderboard{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
