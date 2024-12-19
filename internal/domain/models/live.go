package models

import "time"

type Live struct {
	Timestamp

	Topic       string     `json:"topic"`
	LiveGroupID uint       `json:"liveGroupID"`
	CourseID    uint       `json:"courseID" validate:"required"`
	SubjectID   uint       `json:"subjectID" validate:"required"`
	Type        uint       `json:"type"` //only 2 and 8 : scheduled meeting and recurring meeting, given by zoom
	StartTime   *time.Time `json:"start_time"`
	EndDateTime *time.Time `json:"endDateTime"` //only in recurring live
	Duration    int        `json:"duration"`
	IsLive      *bool      `json:"isLive" gorm:"default:false"` //if only true live can be started
	MeetingID   int        `json:"meetingID"`
	MeetingPwd  string     `json:"password"`
	Email       string     `json:"email"`
	IsFree      *bool      `json:"isFree" gorm:"default:false"`
	IsPackage   *bool      `json:"isPackage" gorm:"default:false"`
	Price       int        `json:"price"`
}
