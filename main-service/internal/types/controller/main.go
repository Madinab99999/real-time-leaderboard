package controller

import "context"

type Controller interface {
	Score
	Leaderboard
}

type Score interface {
	NewScore(context.Context, NewScoreReq) (NewScoreResp, error)
}

type Leaderboard interface {
	GetLeaderboard(context.Context, GetLeaderboardReq) (GetLeaderboardResp, error)
	GetUserRank(context.Context, GetUserRankReq) (GetUserRankResp, error)
	//GetReport(context.Context, GetReportReq) (GetReportResp, error)
}
