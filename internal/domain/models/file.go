package models

type File struct {
	Timestamp
	Title    string
	Type     string
	Url      string `json:"url"`
	IsActive *bool  `gorm:"default:false"`
}
