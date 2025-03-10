package database

type SendNotificationReq interface {
	GetUserID() string
	GetGameName() string
}

type SendNotificationResp interface{}
