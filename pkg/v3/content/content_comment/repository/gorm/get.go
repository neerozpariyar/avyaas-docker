package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetCommentByID(id uint) (models.Comment, error) {
	var comment models.Comment

	// Retrieve the comment from the database based on given id
	err := repo.db.Where("id = ?", id).First(&comment).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Comment{}, fmt.Errorf("comment with comment id: '%d' not found", id)
		}

		return models.Comment{}, err
	}

	return comment, nil
}
