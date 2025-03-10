package leaderboard

import (
	"log/slog"
	"main_service/internal/types/controller"
)

type Leaderboard struct {
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(logger *slog.Logger, ctrl controller.Controller) *Leaderboard {
	return &Leaderboard{
		logger: logger,
		ctrl:   ctrl,
	}
}
