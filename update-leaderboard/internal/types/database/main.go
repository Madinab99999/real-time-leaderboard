package database

import "context"

type Database interface {
	Leaderboard
}

type Leaderboard interface {
	StartUpdate(ctx context.Context) error
	UpdateLeaderboard(ctx context.Context, req UpdateLeaderboardReq) (UpdateLeaderboardResp, error)
	SendMessages(ctx context.Context, userID, game string, newScore float64) error
}
