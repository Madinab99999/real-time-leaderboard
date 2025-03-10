package score

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"main_service/internal/types/database"
)

func (m *Score) AddScore(ctx context.Context, req database.AddScoreReq) (database.AddScoreResp, error) {
	log := m.logger.With(slog.String("handler", "AddScore"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}
	event := newAddScoreEvent(req.GetUserID(), req.GetGameName(), req.GetScore())
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.ErrorContext(ctx, "failed to marshal event", slog.Any("error", err))
		return nil, err
	}

	err = m.db.Publish(ctx, m.conf.RedisChannel, eventJSON).Err()
	if err != nil {
		log.ErrorContext(ctx, "failed to publish message", slog.Any("error", err))
		return nil, err
	}

	log.InfoContext(
		ctx,
		"success publish message to 1 subscribe",
	)
	return true, nil
}

type AddScoreEvent struct {
	UserID string  `json:"user_id"`
	Game   string  `json:"game"`
	Score  float64 `json:"score"`
}

func newAddScoreEvent(user_id string, game string, score float64) *AddScoreEvent {
	return &AddScoreEvent{
		UserID: user_id,
		Game:   game,
		Score:  score,
	}
}
