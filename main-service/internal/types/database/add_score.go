package database

type AddScoreReq interface {
	GetUserID() string
	GetGameName() string
	GetScore() float64
}

type AddScoreResp interface {
}
