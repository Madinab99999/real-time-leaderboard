package score

import (
	"context"
	"fmt"
	"log/slog"
	"main_service/internal/types/controller"
)

func (r *Score) NewScore(ctx context.Context, req controller.NewScoreReq) (controller.NewScoreResp, error) {
	log := r.logger.With(slog.String("handler", "NewScore"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	dbReq := newAddScoreDBReq(req.GetUserID(), req.GetGameName(), req.GetScore())
	dbResp, err := r.db.AddScore(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "db request failed", slog.Any("error", err))
		return nil, fmt.Errorf("db request failed %w", err)
	}
	if dbResp == nil {
		return nil, nil
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}

type addScoreDBReq struct {
	userId string
	game   string
	score  float64
}

func newAddScoreDBReq(userId string, game string, score float64) *addScoreDBReq {
	return &addScoreDBReq{
		userId: userId,
		game:   game,
		score:  score,
	}
}

func (req *addScoreDBReq) GetUserID() string {
	return req.userId
}

func (req *addScoreDBReq) GetGameName() string {
	return req.game
}

func (req *addScoreDBReq) GetScore() float64 {
	return req.score
}
