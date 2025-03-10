package controller

type NewScoreReq interface {
	GetUserID() string
	GetGameName() string
	GetScore() float64
}

type NewScoreResp interface {
}
