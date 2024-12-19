package presenter

import (
	"avyaas/internal/domain/models"
	"time"
)

type ReplyCreateUpdateRequest struct {
	ID           uint   `json:"id"`
	Reply        string `json:"reply" validate:"required"`
	DiscussionID uint   `json:"discussionID" validate:"required"`
	CourseID     uint   `json:"courseID" validate:"required"`
	CreatedBy    uint   `json:"createdBy"`
}
type ReplyListResponse struct {
	ID           uint          `json:"id"`
	DiscussionID uint          `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"discussionID"`
	CreatedAt    time.Time     `json:"createdAt"`
	Discussion   Discussion    ` json:"-"`
	Reply        string        `json:"reply"`
	CourseID     uint          `json:"courseID"`
	Course       models.Course `json:"-"`
	User         interface{}   `json:"user"`
}
type ReplyListRequest struct {
	PageSize     int
	Page         int
	DiscussionID uint
	Search       string
}
