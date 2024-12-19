package models

import "time"

type LiveGroup struct {
	Timestamp

	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description"` //shown before buying
	CourseID    uint       `json:"courseID" validate:"required"`
	StartDate   *time.Time `json:"startDate" validate:"required"`

	// for package data
	IsPackage *bool `json:"isPackage" gorm:"default:false"`
	Price     int   `json:"price"`
	Period    int   `json:"period"`

	// IsPremium   *bool  `json:"isPremium" gorm:"default:false"`
	// Amount      uint   `json:"amount" validate:"required"` //shown before buying
	// ParticipantLimit uint   `json:"participantLimit"`
}
