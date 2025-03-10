package leaderboard

import (
	"context"
	"fmt"
	"log/slog"
	"notification/internal/types/database"

	"github.com/redis/go-redis/v9"
)

func (m *Leaderboard) SendNotification(ctx context.Context, req database.SendNotificationReq) (database.SendNotificationResp, error) {
	log := m.logger.With(slog.String("handler", "SendNotification"))

	gameLeaderboardKey := fmt.Sprintf("leaderboard:game:%s", req.GetGameName())

	notification := fmt.Sprintf("User %s has a new score in game %s", req.GetUserID(), req.GetGameName())
	log.InfoContext(ctx, "notification send for all users", slog.Any("notification", notification))

	gameRank, err := m.db.ZRevRank(ctx, gameLeaderboardKey, req.GetUserID()).Result()
	if err != nil && err != redis.Nil {
		log.ErrorContext(ctx, "failed to get new rank in game", slog.Any("error", err))
		return nil, err
	}
	gameRank = gameRank + 1
	notificationGame := fmt.Sprintf("User %s has new position %d in leaderboard of game %s", req.GetUserID(), gameRank, req.GetGameName())
	log.InfoContext(ctx, "notification send for all users", slog.Any("notification", notificationGame))

	globalKey := "leaderboard:global"
	globalRank, err := m.db.ZRevRank(ctx, globalKey, req.GetUserID()).Result()
	if err != nil && err != redis.Nil {
		log.ErrorContext(ctx, "failed to get global rank", slog.Any("error", err))
		return nil, err
	}
	globalRank = globalRank + 1
	notificationGlobal := fmt.Sprintf("User %s has new position %d in global leaderboard", req.GetUserID(), globalRank)
	log.InfoContext(ctx, "notification send for all users", slog.Any("notification", notificationGlobal))
	return nil, nil
}
