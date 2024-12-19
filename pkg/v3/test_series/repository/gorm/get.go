package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetTestSeriesByID(id uint) (*models.TestSeries, error) {
	var testSeries *models.TestSeries

	err := repo.db.Where("id = ?", id).First(&testSeries).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("test series with id: '%d' not found", id)
		}

		return nil, err
	}

	return testSeries, nil
}

func (repo *Repository) GetTestSeriesByTitle(title string) (*models.TestSeries, error) {
	var testSeries *models.TestSeries
	err := repo.db.Where("title = ?", title).First(&testSeries).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("test series with title: '%s' not found", title)
		}
		return nil, err
	}

	return testSeries, nil
}
