package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) GetConsumedContentCount(userID, courseID uint) (int64, error) {
	var count int64

	err := repo.db.Model(&models.StudentContent{}).
		Where("user_id = ? AND course_id=? AND has_completed = true", userID, courseID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}
