package models

type Note struct {
	Timestamp

	Title       string  `json:"title"`
	Description string  `json:"description"`
	File        string  `json:"file"`
	ContentID   uint    `json:"contentID"`
	Content     Content `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
}
