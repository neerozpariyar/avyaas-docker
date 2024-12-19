package presenter

import "mime/multipart"

type NoteCreateUpdateRequest struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title" validate:"required"`
	Description string                `json:"description"`
	File        *multipart.FileHeader `json:"file"`
	ContentID   uint                  `json:"contentID" validate:"required"`
}

type ContentNoteCreateUpdateRequest struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	File        *multipart.FileHeader `json:"file"`
	ContentID   uint                  `json:"contentID"`
}

type NoteListRequest struct {
	ContentID uint
	Search    string
	Page      int
	PageSize  int
}
