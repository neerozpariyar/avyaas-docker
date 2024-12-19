package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) UpdatePayment(paymentID uint, request *models.Payment) error {
	return repo.db.Where("id = ?", request.ID).Updates(&request).Error
}
