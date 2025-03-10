package governor

import (
	"context"
	"log/slog"
	"main_service/internal/governor/leaderboard"
	"main_service/internal/governor/score"

	"main_service/internal/config"
	"main_service/internal/types/database"
)

type Governor struct {
	*score.Score
	*leaderboard.Leaderboard
}

func New(conf *config.Config) *Governor {
	return &Governor{
		Score:       new(score.Score),
		Leaderboard: new(leaderboard.Leaderboard),
	}
}

func (g *Governor) Config(ctx context.Context, conf *config.Config, logger *slog.Logger, db database.Database) {
	*g.Score = *score.New(conf, logger.With(slog.String("component", "score")), db)
	*g.Leaderboard = *leaderboard.New(conf, logger.With(slog.String("component", "leaderboard")), db)
}
