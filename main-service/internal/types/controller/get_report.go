package controller

type GetReportReq interface {
	GetTopCount() string
	GetGameName() string
}

type GetReportResp interface {
	GetList() []PlayersResp
}

type PlayersResp interface {
	GetRank() int
	GetUserID() string
	GetGameName() string
	GetScore() float32
}
