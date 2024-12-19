package models

type Recording struct {
	Timestamp
	Title       string `json:"title"`
	Description string `json:"description"`
	LiveID      uint   `json:"liveID"`
	Views       int    `json:"views" gorm:"default:0"`
	Url         string `json:"url"`
	Length      uint   `json:"length"`
	Position    uint   `json:"-"`
}
