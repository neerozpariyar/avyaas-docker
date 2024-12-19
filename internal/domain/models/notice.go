package models

type Notice struct {
	Timestamp
	CreatedBy   uint   `json:"-"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
	CourseID    uint   `json:"courseID"`
}
