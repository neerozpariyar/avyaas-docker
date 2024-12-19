package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetLiveGroupByID(id uint) (models.LiveGroup, error) {
	var liveGroup models.LiveGroup

	// Retrieve the liveGroup from the database based on given id
	err := repo.db.Where("id = ?", id).First(&liveGroup).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.LiveGroup{}, fmt.Errorf("liveGroup with liveGroup id: '%d' not found", id)
		}

		return models.LiveGroup{}, err
	}

	return liveGroup, nil
}

func (repo *Repository) GetLiveGroupByTitle(title string) (models.LiveGroup, error) {
	var liveGroup models.LiveGroup

	// Retrieve the liveGroup from the database based on given id
	err := repo.db.Where("title = ?", title).First(&liveGroup).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.LiveGroup{}, fmt.Errorf("liveGroup with title: '%s' not found", title)
		}

		return models.LiveGroup{}, err
	}

	return liveGroup, nil
}
