package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) UpdatePoll(poll models.Poll) error {
	return repo.db.Where("id = ?", poll.ID).Updates(&poll).Error
}
