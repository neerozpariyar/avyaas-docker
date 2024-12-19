package models

type Discussion struct {
	Timestamp
	Title      string  `json:"title"`
	Query      string  `json:"query"`
	VoteCount  uint    `json:"voteCount"`
	Votes      []Vote  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	ReplyCount uint    `json:"replyCount"`
	Replies    []Reply `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	Views      uint    `json:"views"`
	UserID     uint    `json:"-"`
	CourseID   uint    `json:"courseID"`
	Course     Course  `json:"-"`
	SubjectID  uint    `json:"subjectID"`
	Subject    Subject `json:"-"`
}
