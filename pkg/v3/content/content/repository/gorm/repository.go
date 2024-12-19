package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	db       *gorm.DB
	noteRepo interfaces.NoteRepository
}

func New(db *gorm.DB, noteRepo interfaces.NoteRepository) *Repository {
	return &Repository{
		db:       db,
		noteRepo: noteRepo,
	}
}
