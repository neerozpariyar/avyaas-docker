package models

import "time"

// TestResult represents a database model for storing a student's test result
type TestResult struct {
	Timestamp

	UserID           uint
	CourseID         uint
	TestID           uint
	Type             string
	StartTime        *time.Time
	EndTime          *time.Time
	Score            float64
	Percentage       float64
	Rank             int
	TotalAttempted   int
	TotalUnattempted int
	TotalCorrect     int
	TotalWrong       int
}

// TestResponse represents a database model for storing student's each question response of the test
type TestResponse struct {
	Timestamp

	UserID     uint
	TestID     uint
	QuestionID uint
	AnswerID   uint
}
