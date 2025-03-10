package database

type UpdateLeaderboardReq interface {
	GetUserID() string
	GetGameName() string
	GetScore() float64
}

type UpdateLeaderboardResp interface{}
