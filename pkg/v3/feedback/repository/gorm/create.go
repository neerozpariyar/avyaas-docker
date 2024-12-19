package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) CreateFeedback(data models.Feedback) error {
	transaction := repo.db.Begin()

	var courseTitle string
	err := repo.db.Model(&models.Course{}).Select("title").Where("id = ?", data.CourseID).Scan(&courseTitle).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	var existingFeedbacks []models.Feedback
	err = repo.db.Model(&models.Feedback{}).Where("course_id = ? AND user_id = ?", data.CourseID, data.UserID).Find(&existingFeedbacks).Error
	if err == nil && len(existingFeedbacks) != 0 {
		transaction.Rollback()
		return fmt.Errorf("feedback for the given user and course already exists")
	}

	err = transaction.Create(&models.Feedback{
		UserID:      data.UserID,
		Rating:      data.Rating,
		Description: data.Description,
		CourseID:    data.CourseID,
		CourseTitle: courseTitle,
	}).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
