package models

type FCMToken struct {
	Timestamp
	Token      string `json:"token" form:"token"`
	UserID     uint   `json:"user" form:"user"`
	Registered bool   // if token is registered to firebase server
}
