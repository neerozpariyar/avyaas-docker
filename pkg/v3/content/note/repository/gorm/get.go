package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetNoteByID(id uint) (models.Note, error) {
	var note models.Note

	// Retrieve the note from the database based on given id
	err := repo.db.Where("id = ?", id).First(&note).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Note{}, fmt.Errorf("note with note id: '%d' not found", id)
		}

		return models.Note{}, err
	}

	return note, nil
}
