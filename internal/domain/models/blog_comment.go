package models

type BlogComment struct {
	Timestamp
	BlogID  uint   `json:"blogID"`
	UserID  uint   `json:"userID"`
	Comment string `json:"comment"`
}
