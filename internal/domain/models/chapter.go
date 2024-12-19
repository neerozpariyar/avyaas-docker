package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
Chapter represents a database model for storing chapter data. A chapter is an entity that is within
a Unit.
*/
type Chapter struct {
	Timestamp

	Title string `json:"title" validate:"required"`
	// UnitID uint   `json:"unitID" validate:"required"`
	// Unit   Unit   `json:"-"`
	// Contents []Content `gorm:"many2many:chapter_contents;constraint:OnDelete:CASCADE"`
	Position uint `json:"position"`
}

// AfterDelete Hook deletes the nested data along with assosciations of it.

func (m *Chapter) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("id = ?", m.ID).Delete(&Content{})
	return
}

func (m *Chapter) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Model(&m).Association("Contents").Clear()
	return
}
