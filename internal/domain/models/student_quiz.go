package models

// StudentTest represents a database model for storing information about student and test
type StudentTest struct {
	UserID      uint
	CourseID    uint
	TestID      uint
	HasAttended bool `gorm:"default:false"`
	Attempt     uint `gorm:"default:0"`
}
