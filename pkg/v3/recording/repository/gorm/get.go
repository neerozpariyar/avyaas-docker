package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetRecordingByID(id uint) (models.Recording, error) {
	var recording models.Recording

	// Retrieve the recording from the database based on given id
	err := repo.db.Where("id = ?", id).First(&recording).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Recording{}, fmt.Errorf("recording with recording id: '%d' not found", id)
		}

		return models.Recording{}, err
	}

	return recording, nil
}
