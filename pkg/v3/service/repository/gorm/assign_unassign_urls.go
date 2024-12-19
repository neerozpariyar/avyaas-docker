package gorm

import (
	"avyaas/internal/domain/models"

	"gorm.io/gorm"
)

func (repo *Repository) CreateServiceUrls(serviceID uint, urlIDs []uint, transaction *gorm.DB) error {
	for _, urlID := range urlIDs {
		err := transaction.Create(&models.ServiceUrl{
			ServiceID: serviceID,
			UrlID:     urlID,
		}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) DeleteServiceUrls(serviceID uint, urlIDs []uint, transaction *gorm.DB) error {
	for _, urlID := range urlIDs {
		err := transaction.Where("service_id = ? AND url_id = ?", serviceID, urlID).Delete(&models.ServiceUrl{}).Error
		if err != nil {
			return err
		}
	}

	return nil
}
