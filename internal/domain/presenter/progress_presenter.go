package presenter

import "time"

type ProgressPresenter struct {
	ID       uint `json:"id"`
	UserID   uint `json:"userID" `
	CourseID uint `json:"courseID" validate:"required"`
	// ContentID       uint      `json:"contentID"`
	ConsumedContent uint      `json:"consumedContent"`
	LastUpdated     time.Time `json:"lastUpdated"`
}
