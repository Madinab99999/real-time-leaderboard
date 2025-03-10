package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
)

func (m *Leaderboard) StartUpdate(ctx context.Context) error {
	log := m.logger.With(slog.String("handler", "StartUpdate"))
	log.InfoContext(ctx, "sucсess leaderboard service started")
	sub := m.db.Subscribe(ctx, m.conf.RedisChannelUpdate)
	defer sub.Close()

	if _, err := sub.Receive(ctx); err != nil {
		return fmt.Errorf("failed to subscribe to channel %s: %w", m.conf.RedisChannelUpdate, err)
	}

	log.InfoContext(ctx, "successfully subscribed to channel", slog.String("channel", m.conf.RedisChannelUpdate))

	ch := sub.Channel()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				log.WarnContext(ctx, "redis subscription channel closed")
				return nil
			}
			var scoreUpdate UpdateLeaderboardDBReq
			if err := json.Unmarshal([]byte(msg.Payload), &scoreUpdate); err != nil {
				log.ErrorContext(ctx, "failed to parse message payload", slog.Any("error", err))
				continue
			}
			if _, err := m.UpdateLeaderboard(ctx, &scoreUpdate); err != nil {
				log.ErrorContext(ctx, "error processing score update", slog.Any("error", err))
			}
		case <-ctx.Done():
			log.InfoContext(ctx, "stopping leaderboard update service")
			return ctx.Err()
		}
	}
}

type UpdateLeaderboardDBReq struct {
	UserID string  `json:"user_id"`
	Game   string  `json:"game"`
	Score  float64 `json:"score"`
}

func newUpdateLeaderboardDBReq(userId string, game string, score float64) *UpdateLeaderboardDBReq {
	return &UpdateLeaderboardDBReq{
		UserID: userId,
		Game:   game,
		Score:  score,
	}
}

func (req *UpdateLeaderboardDBReq) GetUserID() string {
	return req.UserID
}

func (req *UpdateLeaderboardDBReq) GetGameName() string {
	return req.Game
}

func (req *UpdateLeaderboardDBReq) GetScore() float64 {
	return req.Score
}
