package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"errors"
)

func (repo *Repository) CreateComment(data presenter.CommentCreateUpdateRequest) error {
	transaction := repo.db.Begin()
	// Check if the ContentID has an association with the user
	var count int64

	err := transaction.Model(&models.StudentContent{}).Where("content_id = ? AND user_id = ?", data.ContentID, data.CreatedBy).Count(&count).Error
	if count == 0 {
		transaction.Rollback()
		return errors.New("user has no association with the content")
	}

	if err != nil {
		transaction.Rollback()
		return err
	}

	err = transaction.Create(&models.Comment{
		Comment:   data.Comment,
		ContentID: data.ContentID,
		UserID:    data.CreatedBy,
	}).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
