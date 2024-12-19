package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

func (repo *Repository) GetContentDetails(id uint) (*presenter.ContentDetailResponse, error) {
	content, err := repo.GetContentByID(id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	var contentDetails *presenter.ContentDetailResponse

	// Convert JSON data to a models.Chapter instance
	err = json.Unmarshal(data, &contentDetails)
	if err != nil {
		return nil, err
	}

	var note *models.Note

	err = repo.db.Model(models.Note{}).Where("content_id = ?", id).Scan(&note).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	contentDetails.Note = note

	return contentDetails, nil
}
