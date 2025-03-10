package leaderboard

import (
	"log/slog"
	"main_service/internal/config"
	"main_service/internal/types/database"
)

type Leaderboard struct {
	conf   *config.Config
	logger *slog.Logger
	db     database.Database
}

func New(conf *config.Config, logger *slog.Logger, db database.Database) *Leaderboard {
	return &Leaderboard{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
