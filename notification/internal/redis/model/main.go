package model

import (
	"log/slog"

	"notification/internal/config"
	"notification/internal/redis/model/leaderboard"

	"github.com/redis/go-redis/v9"
)

type Model struct {
	*leaderboard.Leaderboard
}

func New(conf *config.RedisConfig, logger *slog.Logger, db *redis.Client) *Model {
	return &Model{
		Leaderboard: leaderboard.New(conf, logger.With(slog.String("component", "notification")), db),
	}
}
