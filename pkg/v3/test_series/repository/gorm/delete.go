package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) DeleteTestSeries(id uint) error {
	return repo.db.Unscoped().Where("id = ?", id).Delete(&models.TestSeries{}).Error
}
