package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetUnitByID(id uint) (models.Unit, error) {
	var unit models.Unit

	// Retrieve the unit from the database based on given id
	err := repo.db.Where("id = ?", id).First(&unit).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Unit{}, fmt.Errorf("unit with unit id: '%d' not found", id)
		}

		return models.Unit{}, err
	}

	return unit, nil
}

func (repo *Repository) GetChaptersByUnitID(id uint) ([]uint, error) {
	var chaptersID []uint
	err := repo.db.Select("chapter_id").Model(&models.UnitChapterContent{}).Where("unit_id = ?", id).Find(&chaptersID).Error

	if err != nil {
		return nil, err
	}

	return chaptersID, nil
}
