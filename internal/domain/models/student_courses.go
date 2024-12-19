package models

import "time"

type StudentCourse struct {
	Timestamp

	UserID       uint       `json:"userID"`
	CourseID     uint       `json:"courseID"`
	Paid         *bool      `json:"paid" gorm:"default:true"`
	ExpiryDate   *time.Time `json:"expiryDate"`
	Progress     float64    `json:"progress" gorm:"default:0"`
	HasCompleted *bool      `json:"hasCompleted" gorm:"default:false"`
}
