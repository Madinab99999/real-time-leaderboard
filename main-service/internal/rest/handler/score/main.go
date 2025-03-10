package score

import (
	"log/slog"

	"main_service/internal/types/controller"
)

type Score struct {
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(logger *slog.Logger, ctrl controller.Controller) *Score {
	return &Score{
		logger: logger,
		ctrl:   ctrl,
	}
}
