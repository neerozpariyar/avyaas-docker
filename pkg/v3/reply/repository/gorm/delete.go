package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteReply(id uint) error {
	// Perform a hard delete of the reply  with the given ID using the GORM Unscoped method
	return repo.db.Unscoped().Where("id = ?", id).Delete(&models.Reply{}).Error
}
