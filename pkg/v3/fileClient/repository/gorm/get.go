package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) GetObjectsByID(id []uint) ([]models.File, error) {
	var objects []models.File

	if err := repo.db.Debug().Where("id IN (?)", id).Find(&objects).Error; err != nil {
		return nil, err
	}

	return objects, nil
}

func (repo *Repository) GetURLsByID(id []uint) ([]string, error) {
	var urls []string

	if err := repo.db.Debug().Model(&models.File{}).Where("id IN (?)", id).Pluck("url", &urls).Error; err != nil {
		return nil, err
	}

	fmt.Printf("urls: %v\n", urls)
	return urls, nil
}
