package database

import "context"

type Database interface {
	Score
	Leaderboard
}

type Score interface {
	AddScore(context.Context, AddScoreReq) (AddScoreResp, error)
}

type Leaderboard interface {
	GetTopPlayers(context.Context, GetTopPlayersReq) (GetTopPlayersResp, error)
	GetUserRank(context.Context, GetUserRankReq) (GetUserRankResp, error)
	//GetReport(context.Context, GetReportReq) (GetReportResp, error)
}
