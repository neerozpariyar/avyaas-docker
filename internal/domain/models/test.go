package models

import "time"

// TestType represents a database model struct for storing information about a test type.
type TestType struct {
	Timestamp

	Title string `json:"title" validate:"required"`
}

// Test represents a database model struct for storing information about a test.
type Test struct {
	Timestamp

	Title        string
	StartTime    *time.Time
	EndTime      *time.Time
	Duration     int
	ExtraTime    int `gorm:"default:15"`
	Price        int
	CourseID     uint
	Course       Course `json:"-"`
	TestTypeID   int
	TestSeriesID uint
	Status       string // Scheduled, InActive, Expired
	IsPublic     *bool  `gorm:"default:false"`
	IsPremium    *bool  `gorm:"default:false"`
	IsFree       *bool  `gorm:"default:true"`
	IsMock       *bool  `gorm:"default:false"`
	CreatedBy    uint
	QuestionSets []QuestionSet `gorm:"many2many:test_question_sets;" json:"questionSets"`

	// for package data
	IsPackage *bool `gorm:"default:false"`
}

// type Test struct {
// 	Timestamp

// 	Title        string
// 	StartTime    *time.Time
// 	EndTime      *time.Time
// 	Duration     int
// 	ExtraTime    int `gorm:"default:15"`
// 	Price        int
// 	TestTypeID   uint
// 	TestType     TestType `json:"-"`
// 	CourseID     uint
// 	Course       Course `json:"-"`
// 	TestSeriesID uint
// 	IsDraft      *bool `gorm:"default:false"`
// 	IsPremium    *bool `gorm:"default:false"`
// 	IsFree       *bool `gorm:"default:false"`
// 	CreatedBy    uint
// 	QuestionSets []QuestionSet `gorm:"many2many:test_question_sets;" json:"questionSets"`
// }

// TestQuestionSet represents a database model for storing association between test and question set
type TestQuestionSet struct {
	TestID        uint
	QuestionSetID uint
}
