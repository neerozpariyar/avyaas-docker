package models

type PollVote struct {
	Timestamp

	PollID       uint       `json:"pollID"`
	Poll         Poll       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	PollOptionID uint       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"pollOptionID"`
	PollOption   PollOption `json:"-"`
	UserID       uint       `json:"userID"`
}
