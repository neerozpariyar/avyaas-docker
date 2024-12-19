package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) CreateService(data presenter.ServiceCreateUpdateRequest) error {
	var newService *models.Service
	transaction := repo.db.Begin()

	err := transaction.Create(&models.Service{
		Title:       data.Title,
		Description: data.Description,
	}).Scan(&newService).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	// err = repo.CreateServiceUrls(newService.ID, data.UrlIDs, transaction)
	// if err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }

	transaction.Commit()
	return nil
}
