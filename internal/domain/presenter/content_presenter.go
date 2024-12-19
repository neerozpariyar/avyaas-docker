package presenter

import (
	"avyaas/internal/domain/models"

	"mime/multipart"
)

type ContentCreateUpdateRequest struct {
	ID          uint                           `json:"id"`
	Title       string                         `json:"title" validate:"required"`
	Description string                         `json:"description"`
	IsPremium   *bool                          `json:"isPremium" gorm:"default:false" validate:"required"`
	ContentType string                         `json:"contentType" validate:"required"`
	Level       string                         `json:"level" validate:"required"`
	Visibility  *bool                          `json:"visibility" gorm:"default:false" validate:"required"`
	File        *multipart.FileHeader          `json:"file" form:"file" validate:"required"`
	Length      uint                           `json:"length"`
	CreatedBy   uint                           `json:"createdBy"`
	ChapterID   uint                           `json:"chapterID"`
	HasNote     *bool                          `json:"hasNote"`
	Note        ContentNoteCreateUpdateRequest `json:"note"`
}

type ContentListRequest struct {
	Page          int
	PageSize      int
	Search        string
	ContentFilter FilterContentRequest
	UserID        uint
}

type FilterContentRequest struct {
	ChapterID uint `json:"chapterID"  validate:"required"`
	SubjectID uint `json:"subjectID,omitempty" validate:"required"`
	UnitID    uint `json:"unitID,omitempty" validate:"required"`
}
type UpdateContentPositionRequest struct {
	ContentIDs []uint `json:"contentIDs"`
}

type SingleContentResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	IsPremium    bool   `json:"isPremium"`
	ContentType  string `json:"contentType"`
	Length       uint   `json:"length"`
	Paid         *bool  `json:"paid"`
	HasCompleted bool   `json:"hasCompleted"`
}

type ContentListResponse struct {
	Success     bool             `json:"success"`
	CurrentPage int32            `json:"currentPage"`
	TotalPage   int32            `json:"totalPage"`
	Data        []models.Content `json:"data"`
}

type ContentDetailResponse struct {
	ID           uint         `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	IsPremium    *bool        `json:"isPremium"`
	Paid         *bool        `json:"paid"`
	ContentType  string       `json:"contentType"`
	Length       uint         `json:"length"`
	Level        string       `json:"level"`
	Visibility   *bool        `json:"visibility"`
	CourseID     uint         `json:"courseID,omitempty"`
	ChapterID    uint         `json:"chapterID,omitempty"`
	Views        int          `json:"views"`
	Url          string       `json:"url,omitempty"`
	Note         *models.Note `json:"note"`
	IsBookmarked bool         `json:"isBookmarked"`
	BookmarkID   uint         `json:"bookmarkID"`
	Progress     float64      `json:"progress"`
	HasCompleted bool         `json:"hasCompleted"`
}
