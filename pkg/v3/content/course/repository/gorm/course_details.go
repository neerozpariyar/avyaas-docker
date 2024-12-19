package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) GetCourseDetails(id uint) (*models.Course, error) {
	var course *models.Course

	if err := repo.db.Preload("Subjects.Units.Chapters.Contents").First(&course, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return course, nil
}
