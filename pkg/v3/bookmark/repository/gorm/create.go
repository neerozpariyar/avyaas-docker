package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) CreateBookmark(data presenter.BookmarkCreateUpdateRequest) error {
	transaction := repo.db.Begin()

	// var courseTitle string
	// err := repo.db.Model(&models.Course{}).Select("title").Where("id = ?", data.ContentID).Scan(&courseTitle).Error
	// if err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }

	// var questionTitle string
	// err = repo.db.Model(&models.Question{}).Select("title").Where("id = ?", data.QuestionID).Scan(&questionTitle).Error
	// if err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }
	if data.ContentID == 0 && data.QuestionID == 0 {
		return fmt.Errorf("either contentID or questionID must be provided")
	}

	var existingContent models.StudentContent
	var existingBookmarks []models.Bookmark

	if data.ContentID != 0 {
		err := repo.db.Model(&models.StudentContent{}).Where("content_id = ? AND user_id = ?", data.ContentID, data.UserID).First(&existingContent).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				transaction.Rollback()
				return fmt.Errorf("content does not exist for the given user")
			}
			return err
		}

		err = repo.db.Model(&models.Bookmark{}).Where("content_id = ? AND user_id = ?", data.ContentID, data.UserID).Find(&existingBookmarks).Error
		if err == nil && len(existingBookmarks) != 0 {
			transaction.Rollback()
			return fmt.Errorf("bookmark for the given user and content already exists")
		}
	}

	var existingQuestion models.Question

	if data.QuestionID != 0 {
		err := repo.db.Model(&models.Question{}).Where("id = ?", data.QuestionID).First(&existingQuestion).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				transaction.Rollback()
				return fmt.Errorf("question does not exist ")
			}
			return err
		}

		err = repo.db.Model(&models.Bookmark{}).Where("question_id = ? AND user_id = ?", data.QuestionID, data.UserID).Find(&existingBookmarks).Error
		if err == nil && len(existingBookmarks) != 0 {
			transaction.Rollback()
			return fmt.Errorf("bookmark for the given user and question already exists")
		}
	}
	var bookmarkType string

	if data.QuestionID != 0 {
		bookmarkType = "question"
	} else {
		bookmarkType = "content"
	}

	err := transaction.Create(&models.Bookmark{
		UserID:       data.UserID,
		QuestionID:   data.QuestionID,
		ContentID:    data.ContentID,
		BookmarkType: bookmarkType,
	}).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
