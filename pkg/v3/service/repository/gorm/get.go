package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetServiceByID(id uint) (*models.Service, error) {
	var service *models.Service

	err := repo.db.Where("id = ?", id).First(&service).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("service with id: '%d' not found", id)
		}

		return nil, err
	}

	return service, nil
}

func (repo *Repository) GetServiceByTitle(title string) (*models.Service, error) {
	var service *models.Service

	err := repo.db.Where("title = ?", title).First(&service).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("service with title: '%s' not found", title)
		}

		return nil, err
	}

	return service, nil
}

func (repo *Repository) GetUrlByID(id uint) error {
	var urlID uint
	err := repo.db.Table("urls").Select("id").Where("id = ?", id).Scan(&urlID).Error
	if err != nil {
		return err
	}

	if urlID == 0 {
		return fmt.Errorf("url with id: '%d' not found", id)
	}

	return nil
}

func (repo *Repository) GetUrlIDsByServiceID(serviceID uint) ([]uint, error) {
	var urlIDs []uint

	err := repo.db.Model(&models.ServiceUrl{}).Select("url_id").Where("service_id = ?", serviceID).Scan(&urlIDs).Error
	if err != nil {
		return urlIDs, err
	}

	return urlIDs, nil
}
