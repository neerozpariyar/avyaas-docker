package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetContentByID(id uint) (models.Content, error) {
	var content models.Content

	// Retrieve the content from the database based on given id
	err := repo.db.Where("id = ?", id).First(&content).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Content{}, fmt.Errorf("content with content id: '%d' not found", id)
		}

		return models.Content{}, err
	}

	return content, nil
}

func (repo *Repository) CheckStudentContent(userID, contentID uint) (models.StudentContent, error) {
	var content models.StudentContent

	// Retrieve the course from the database based on given courseID
	err := repo.db.Where("user_id = ? And content_id = ?", userID, contentID).First(&content).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.StudentContent{}, fmt.Errorf("course not subscribed for premium contents")
		}

		return models.StudentContent{}, err
	}

	return content, nil
}

func (repo *Repository) GetContentProgressByContentID(contentID, userID uint) (models.StudentContent, error) {
	var content models.StudentContent
	// Retrieve the course from the database based on given courseID
	err := repo.db.Debug().Where("content_id = ? AND user_id=?", contentID, userID).First(&content).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.StudentContent{}, fmt.Errorf("content not found with content id: '%d'", contentID)
		}

		return models.StudentContent{}, err
	}

	return content, nil
}
