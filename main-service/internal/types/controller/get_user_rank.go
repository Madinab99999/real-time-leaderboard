package controller

type GetUserRankReq interface {
	GetUserID() string
}

type GetUserRankResp interface {
	GetGlobalRank() int
	GetGlobalScore() float64
	GetScoreList() []ScoresOfGames
}

type ScoresOfGames interface {
	GetRank() int
	GetGameName() string
	GetScore() float64
}
