package database

type SendMessagesReq interface {
	GetUserID() string
	GetGameName() string
}

type SendMessagesResp interface{}
