package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
Subject represents a database model for storing subject data. A subject is an entity that is within
a Course.
*/
type Subject struct {
	Timestamp

	SubjectID          string               `json:"subjectID"` // slug of subject
	Title              string               `json:"title"`
	Description        string               `json:"description"`
	Thumbnail          string               `json:"thumbnail"`
	Polls              []Poll               `json:"-"`
	UnitChapterContent []UnitChapterContent `gorm:"many2many:subject_unit_chapter_contents;constraint:onDelete:CASCADE"`
}

type SubjectUnitChapterContent struct {
	SubjectID            uint
	Position             uint
	UnitChapterContentID uint
}

func (m *Subject) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("id = ?", m.ID).Delete(&Unit{})
	return
}

func (m *Subject) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Model(&m).Association("UnitChapterContent").Clear()
	return
}

func (m *SubjectUnitChapterContent) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64

	err = tx.Model(&SubjectUnitChapterContent{}).Where("subject_id = ?", m.SubjectID).Count(&count).Error
	m.Position = uint(count) + 1

	return
}
