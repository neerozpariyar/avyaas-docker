package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetChapterByID(id uint) (models.Chapter, error) {
	var chapter models.Chapter

	// Retrieve the chapter from the database based on given id
	err := repo.db.Where("id = ?", id).First(&chapter).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Chapter{}, fmt.Errorf("chapter with chapter id: '%d' not found", id)
		}

		return models.Chapter{}, err
	}

	return chapter, nil
}
