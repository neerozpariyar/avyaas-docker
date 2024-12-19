package gorm

import (
	"avyaas/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

func (repo *Repository) UpdateFeedback(data models.Feedback) error {
	_, err := repo.GetFeedbackByID(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return repo.db.Where("id = ?", data.ID).Updates(&models.Feedback{
		UserID:      data.UserID,
		Rating:      data.Rating,
		Description: data.Description,
		CourseID:    data.CourseID,
	}).Error
}
