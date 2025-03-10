package database

import "context"

type Database interface {
	Leaderboard
}

type Leaderboard interface {
	StartNotification(ctx context.Context) error
	SendNotification(ctx context.Context, req SendNotificationReq) (SendNotificationResp, error)
	SendMessages(ctx context.Context, userID, game string, newScore float64) error
}
