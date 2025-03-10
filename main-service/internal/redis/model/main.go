package model

import (
	"log/slog"

	"main_service/internal/config"
	"main_service/internal/redis/model/leaderboard"
	"main_service/internal/redis/model/score"

	"github.com/redis/go-redis/v9"
)

type Model struct {
	*score.Score
	*leaderboard.Leaderboard
}

func New(conf *config.RedisConfig, logger *slog.Logger, db *redis.Client) *Model {
	return &Model{
		Score:       score.New(conf, logger.With(slog.String("component", "score")), db),
		Leaderboard: leaderboard.New(conf, logger.With(slog.String("component", "leaderboard")), db),
	}
}
