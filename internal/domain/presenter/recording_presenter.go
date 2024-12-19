package presenter

import "mime/multipart"

type RecordingCreateUpdateRequest struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title" validate:"required"`
	Description string                `json:"description"`
	LiveID      uint                  `json:"liveID" validate:"required"`
	Views       int                   `json:"views" gorm:"default:0"`
	File        *multipart.FileHeader `json:"file" validate:"required"`
	Length      uint                  `json:"length" validate:"required"`
}

type RecordingListRequest struct {
	PageSize int
	Page     int
	LiveID   uint
	Search   string
}
