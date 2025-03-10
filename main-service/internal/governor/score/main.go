package score

import (
	"log/slog"
	"main_service/internal/config"
	"main_service/internal/types/database"
)

type Score struct {
	conf   *config.Config
	logger *slog.Logger
	db     database.Database
}

func New(conf *config.Config, logger *slog.Logger, db database.Database) *Score {
	return &Score{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
