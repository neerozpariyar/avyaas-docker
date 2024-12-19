package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateLiveGroup(data models.LiveGroup) error {
	return repo.db.Where("id = ?", data.ID).Updates(&data).Error
}
