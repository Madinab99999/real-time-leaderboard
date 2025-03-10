package leaderboard

import (
	"context"
	"encoding/json"
	"log/slog"
)

func (m *Leaderboard) SubscribeLeaderboardUpdates() {
	ctx := context.Background()
	pubsub := m.db.Subscribe(ctx, m.conf.RedisChannel)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		var event map[string]string
		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			m.logger.ErrorContext(ctx, "failed to unmarshal event", slog.Any("error", err))
			continue
		}

		if event["event"] == "updated" {
			m.PrintLeaderboard()
		}
	}
}
