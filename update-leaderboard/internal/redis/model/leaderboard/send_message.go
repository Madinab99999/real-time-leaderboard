package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"update-leaderboard/internal/types/database"
)

func (m *Leaderboard) SendMessages(ctx context.Context, req database.SendMessagesReq) (database.UpdateLeaderboardResp, error) {
	log := m.logger.With(slog.String("handler", "SendMessages"))

	event := map[string]string{"event": "updated"}
	eventPayload, err := json.Marshal(event)
	if err != nil {
		log.ErrorContext(ctx, "failed to marshal event", slog.Any("error", err))
		return nil, fmt.Errorf("failed to marshal event: %w", err)
	}
	if err := m.db.Publish(ctx, m.conf.RedisChannelConsole, eventPayload).Err(); err != nil {
		log.ErrorContext(ctx, "failed to publish message", slog.Any("error", err))
		return nil, err
	}

	notification := newNotificationEvent(req.GetUserID(), req.GetGameName())

	payload, err := json.Marshal(notification)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal notification: %w", err)
	}

	if err := m.db.Publish(ctx, m.conf.RedisChannelNotification, payload).Err(); err != nil {
		return nil, fmt.Errorf("failed to publish message : %w", err)
	}

	log.InfoContext(ctx, "success publish message of update to 2 subscribe")

	return nil, nil
}

type NotificationEvent struct {
	UserID string `json:"user_id"`
	Game   string `json:"game"`
}

func newNotificationEvent(user_id string, game string) *NotificationEvent {
	return &NotificationEvent{
		UserID: user_id,
		Game:   game,
	}
}
