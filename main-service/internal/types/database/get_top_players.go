package database

type GetTopPlayersReq interface {
	GetTopCount() int
	GetGameName() *string
}

type GetTopPlayersResp interface {
	GetList() []ItemGetTopPlayersResp
}

type ItemGetTopPlayersResp interface {
	GetRank() int
	GetUserID() string
	GetScore() float64
}
