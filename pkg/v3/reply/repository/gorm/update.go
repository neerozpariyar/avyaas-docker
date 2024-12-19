package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateReply(reply models.Reply) error {
	return repo.db.Where("id = ?", reply.ID).Updates(&reply).Error
}
