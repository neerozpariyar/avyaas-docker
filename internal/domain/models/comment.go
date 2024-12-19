package models

type Comment struct {
	Timestamp

	ContentID uint    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"contentID"`
	Content   Content ` json:"-"`
	Comment   string  `json:"comment"`
	UserID    uint    `json:"userID"`
}
