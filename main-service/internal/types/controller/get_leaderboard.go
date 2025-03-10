package controller

type GetLeaderboardReq interface {
	GetTopCount() int
	GetGameName() *string
}

type GetLeaderboardResp interface {
	GetList() []ItemGetLeaderboardResp
}

type ItemGetLeaderboardResp interface {
	GetRank() int
	GetUserID() string
	GetScore() float64
}
