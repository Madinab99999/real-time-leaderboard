package leaderboard

import (
	"context"
	"fmt"
	"log/slog"
	"update-leaderboard/internal/types/database"

	"github.com/redis/go-redis/v9"
)

func (m *Leaderboard) UpdateLeaderboard(ctx context.Context, req database.UpdateLeaderboardReq) (database.UpdateLeaderboardResp, error) {
	log := m.logger.With(slog.String("handler", "UpdateLeaderboard"))

	gameKey := gameKey(req.GetGameName())
	globalKey := "leaderboard:global"

	err := m.db.Watch(ctx, func(tx *redis.Tx) error {
		pipe := tx.Pipeline()

		currentScore, err := tx.ZScore(ctx, gameKey, req.GetUserID()).Result()
		if err != nil && err != redis.Nil {
			return fmt.Errorf("failed to get current score: %w", err)
		}

		newScore := req.GetScore() - currentScore

		if newScore != 0 {
			if err := pipe.ZIncrBy(ctx, globalKey, newScore, req.GetUserID()).Err(); err != nil {
				return fmt.Errorf("failed to update global leaderboard: %w", err)
			}
		}

		log.InfoContext(ctx, "successfully updated user's score in global leaderboard",
			slog.String("user", req.GetUserID()))

		if err := pipe.ZAdd(ctx, gameKey, redis.Z{
			Score:  req.GetScore(),
			Member: req.GetUserID(),
		}).Err(); err != nil {
			return fmt.Errorf("failed to update game leaderboard: %w", err)
		}

		pipe.SAdd(ctx, "games", req.GetGameName()).Err()
		if err != nil && err != redis.Nil {
			return fmt.Errorf("failed to add game to set: %w", err)
		}

		log.InfoContext(ctx, "successfully updated user's score in leaderboard of game",
			slog.String("user", req.GetUserID()),
			slog.String("game", req.GetGameName()))
		mes := newSendMessagesDBReq(req.GetUserID(), req.GetGameName())
		if _, err := m.SendMessages(ctx, mes); err != nil {
			log.ErrorContext(ctx, "failed to publish message", slog.Any("error", err))
			return err
		}

		if _, err := pipe.Exec(ctx); err != nil {
			return fmt.Errorf("failed to execute transaction: %w", err)
		}

		return nil
	}, gameKey)
	return nil, err
}

func gameKey(game string) string {
	return fmt.Sprintf("leaderboard:game:%s", game)
}

type SendMessagesDBReq struct {
	UserID string `json:"user_id"`
	Game   string `json:"game"`
}

func newSendMessagesDBReq(userId string, game string) *SendMessagesDBReq {
	return &SendMessagesDBReq{
		UserID: userId,
		Game:   game,
	}
}

func (req *SendMessagesDBReq) GetUserID() string {
	return req.UserID
}

func (req *SendMessagesDBReq) GetGameName() string {
	return req.Game
}
