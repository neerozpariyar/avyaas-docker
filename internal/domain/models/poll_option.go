package models

type PollOption struct {
	Timestamp

	PollID uint       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"pollID"`
	Poll   Poll       `json:"-"`
	Option string     `json:"option"`
	Votes  []PollVote `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	UserID uint       `json:"userID"`
}
