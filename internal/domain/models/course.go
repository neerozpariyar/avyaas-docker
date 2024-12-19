package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
Course represents a database model for storing course data. A course is an entity that is within a
CourseGroup.
*/
type Course struct {
	Timestamp

	CourseID     string        `json:"courseID"` // slug of course
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Available    *bool         `json:"available" gorm:"default:false"`
	Thumbnail    string        `json:"thumbnail"`
	Subjects     []Subject     `json:"-" gorm:"many2many:course_subjects;constraint:OnDelete:CASCADE"`
	Packages     []Package     `json:"-"`
	Tests        []Test        `json:"-"`
	QuestionSets []QuestionSet `json:"-"`
	Polls        []Poll        `json:"-"`
}

// AfterDelete Hook deletes the nested data along with assosciations of it.

func (m *Course) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("id = ?", m.ID).Delete(&Subject{})
	return
}

func (m *Course) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Model(&m).Association("Subjects").Clear()
	return
}
