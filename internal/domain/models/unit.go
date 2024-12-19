package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
Unit represents database model for storing unit data. A unit is an entity that is within a Subject.
*/
type Unit struct {
	Timestamp

	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	// SubjectID   uint      `json:"subjectID"`
	// Subject     Subject   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	Position uint `json:"-"`
	// Chapters []Chapter `json:"-" gorm:"many2many:unit_chapters;constraint:OnDelete:CASCADE"`
}

type UnitChapterContent struct {
	ID        uint `gorm:"primaryKey"`
	UnitID    uint
	ChapterID uint
	ContentID uint
}

func (m *Unit) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("id = ?", m.ID).Delete(&Chapter{})
	return
}

func (m *Unit) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Model(&m).Association("Chapters").Clear()
	return
}
