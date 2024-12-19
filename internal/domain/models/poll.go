package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Poll struct {
	Timestamp

	CourseID  uint         `json:"courseID"`
	Course    Course       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	Options   []PollOption `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	Question  string       `json:"question"`
	SubjectID uint         `json:"subjectID"`
	Subject   Subject      `json:"-"`
	UserID    uint         `json:"userID"`
}

func (m *Poll) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("poll_id = ?", m.ID).Delete(&PollOption{})
	return
}
