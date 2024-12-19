package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) CreateReferral(data models.Referral) error {
	return repo.db.Create(&data).Error
}
