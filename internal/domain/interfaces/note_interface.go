package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type NoteUsecase interface {
	CreateNote(data presenter.NoteCreateUpdateRequest) map[string]string
	ListNote(data presenter.NoteListRequest) ([]models.Note, int, error)
	UpdateNote(data presenter.NoteCreateUpdateRequest) map[string]string
	DeleteNote(id uint) error
}

type NoteRepository interface {
	GetNoteByID(id uint) (models.Note, error)

	CreateNote(data presenter.NoteCreateUpdateRequest) error
	ListNote(data presenter.NoteListRequest) ([]models.Note, float64, error)
	UpdateNote(data presenter.NoteCreateUpdateRequest) error
	DeleteNote(id uint) error
}
