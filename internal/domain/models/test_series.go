package models

import "time"

type TestSeries struct {
	Timestamp

	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description"`
	NoOfTests   int        `json:"noOfTests" validate:"required"`
	CourseID    uint       `json:"courseID" validate:"required"`
	StartDate   *time.Time `json:"startDate" validate:"required"`

	// for package data
	IsPackage *bool `json:"isPackage" gorm:"default:false"`
	Price     int   `json:"price"`
	Period    int   `json:"period"`
}
