package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
)

func (m *Leaderboard) StartNotification(ctx context.Context) error {
	log := m.logger.With(slog.String("handler", "StartNotification"))
	log.InfoContext(ctx, "sucсess notification service started")
	sub := m.db.Subscribe(ctx, m.conf.RedisChannel)
	defer sub.Close()

	if _, err := sub.Receive(ctx); err != nil {
		return fmt.Errorf("failed to subscribe to channel %s: %w", m.conf.RedisChannel, err)
	}

	log.InfoContext(ctx, "successfully subscribed to channel")

	ch := sub.Channel()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				log.WarnContext(ctx, "redis subscription channel closed")
				return nil
			}
			var scoreUpdate SendNotificationDBReq
			if err := json.Unmarshal([]byte(msg.Payload), &scoreUpdate); err != nil {
				log.ErrorContext(ctx, "failed to parse message payload", slog.Any("error", err))
				continue
			}
			if _, err := m.SendNotification(ctx, &scoreUpdate); err != nil {
				log.ErrorContext(ctx, "error processing send notification", slog.Any("error", err))
			}
		case <-ctx.Done():
			log.InfoContext(ctx, "stopping notification service")
			return ctx.Err()
		}
	}
}

type SendNotificationDBReq struct {
	UserID string `json:"user_id"`
	Game   string `json:"game"`
}

func newSendNotificationDBReq(userId string, game string) *SendNotificationDBReq {
	return &SendNotificationDBReq{
		UserID: userId,
		Game:   game,
	}
}

func (req *SendNotificationDBReq) GetUserID() string {
	return req.UserID
}

func (req *SendNotificationDBReq) GetGameName() string {
	return req.Game
}
