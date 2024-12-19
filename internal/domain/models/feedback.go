package models

type Feedback struct {
	Timestamp

	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	UserID      uint    `json:"userID"`
	CourseID    uint    `json:"courseID"`
	CourseTitle string  `json:"courseTitle"`
}
