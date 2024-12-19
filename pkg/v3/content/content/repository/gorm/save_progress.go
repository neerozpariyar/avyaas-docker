package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) SaveOrUpdateProgress(progress *models.StudentCourse) error {
	var existingProgress []models.StudentCourse
	err := repo.db.Model(&models.StudentCourse{}).Where("course_id = ?", progress.CourseID).Find(&existingProgress).Error
	if err == nil && len(existingProgress) != 0 {
		return repo.db.Model(&models.StudentCourse{}).
			Where("user_id = ? AND course_id = ?", progress.UserID, progress.CourseID).
			Updates(progress).Error
	} else {
		return repo.db.Create(&progress).Error

	}

}
