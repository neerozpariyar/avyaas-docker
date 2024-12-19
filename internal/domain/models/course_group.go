package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
CourseGroup represents a database model for storing course group data. It represents a category that
a course is listed under.
*/
type CourseGroup struct {
	Timestamp

	GroupID     string   `json:"groupID"` // slug of course group
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Courses     []Course `gorm:"many2many:course_group_courses;constraint:OnDelete:CASCADE;" json:"-"`
}

// AfterDelete Hook deletes the nested data along with assosciations of it.
func (m *CourseGroup) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("id = ?", m.ID).Delete(&Course{})
	return
}

func (m *CourseGroup) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Model(&m).Association("Courses").Clear()
	return
}
