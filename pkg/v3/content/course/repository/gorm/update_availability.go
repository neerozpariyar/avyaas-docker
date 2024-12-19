package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) UpdateAvailability(course models.Course) error {
	return repo.db.Debug().Model(&models.Course{}).Where("id=?", course.ID).Update("available", !*course.Available).Error
}
