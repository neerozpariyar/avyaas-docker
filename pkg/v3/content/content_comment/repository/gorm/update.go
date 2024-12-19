package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateComment(data models.Comment) error {
	return repo.db.Where("id = ?", data.ID).Updates(&data).Error
}
