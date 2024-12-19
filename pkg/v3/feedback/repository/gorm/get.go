package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetFeedbackByID(id uint) (models.Feedback, error) {
	var feedback models.Feedback

	// Retrieve the feedback from the database based on given id
	err := repo.db.Where("id = ?", id).First(&feedback).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Feedback{}, fmt.Errorf("feedback with feedback id: '%d' not found", id)
		}

		return models.Feedback{}, err
	}

	return feedback, nil
}
