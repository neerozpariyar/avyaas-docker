package presenter

import "time"

type CommentCreateUpdateRequest struct {
	ID        uint   `json:"id"`
	Comment   string `json:"comment" validate:"required"`
	ContentID uint   `json:"contentID" validate:"required"`
	// CourseID  uint   `json:"courseID" validate:"required"`
	CreatedBy uint `json:"createdBy"`
}

type CommentListRequest struct {
	PageSize  int
	Page      int
	ContentID uint
	Search    string
}

type CommentListResponse struct {
	ID        uint        `json:"id"`
	ContentID uint        `json:"contentID"`
	Comment   string      `json:"comment"`
	UpdatedAt time.Time   `json:"updatedAt"`
	CreatedBy interface{} `json:"createdBy"`
}
