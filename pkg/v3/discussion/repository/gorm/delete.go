package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteDiscussion(id uint) error {
	// Perform a hard delete of the discussion  with the given ID using the GORM Unscoped method
	return repo.db.Unscoped().Where("id = ?", id).Delete(&models.Discussion{}).Error
}
