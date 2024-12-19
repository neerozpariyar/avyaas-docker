package presenter

type ContentProgressPresenter struct {
	ID     uint `json:"id"`
	UserID uint `json:"userID" `
	// CourseID        uint `json:"courseID" `
	// ContentID       uint `json:"contentID" validate:"required"`
	ElapsedDuration uint `json:"elapsedDuration"`
	TotalDuration   uint `json:"totalDuration"`
}
