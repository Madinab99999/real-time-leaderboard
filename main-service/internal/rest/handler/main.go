package handler

import (
	"log/slog"

	"main_service/internal/rest/handler/leaderboard"
	"main_service/internal/rest/handler/score"
	"main_service/internal/types/controller"
)

type Handler struct {
	*score.Score
	*leaderboard.Leaderboard
}

func New(logger *slog.Logger, ctrl controller.Controller) *Handler {
	return &Handler{
		Score:       score.New(logger.With(slog.String("component", "score")), ctrl),
		Leaderboard: leaderboard.New(logger.With(slog.String("component", "leaderboard")), ctrl),
	}
}
