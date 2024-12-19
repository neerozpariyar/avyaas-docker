package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetBookmarkByID(id uint) (models.Bookmark, error) {
	var bookmark models.Bookmark

	// Retrieve the bookmark from the database based on given id
	err := repo.db.Where("id = ?", id).First(&bookmark).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Bookmark{}, fmt.Errorf("bookmark with bookmark id: '%d' not found", id)
		}

		return models.Bookmark{}, err
	}

	return bookmark, nil
}
func (repo *Repository) GetBookmarkTypeByID(id uint) (string, error) {

	var bookmark models.Bookmark

	// Retrieve the bookmark type from the database based on the given id
	err := repo.db.Where("id = ?", id).First(&bookmark).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Bookmark{}.BookmarkType, fmt.Errorf("bookmark type with id: '%d' not found", id)
		}

		return models.Bookmark{}.BookmarkType, err
	}

	return bookmark.BookmarkType, nil
}
func (repo *Repository) GetBookmarkTitleByID(id uint) (string, error) {
	var bookmark models.Bookmark
	// Retrieve the bookmark from the database based on the given id
	err := repo.db.Where("id = ?", id).First(&bookmark).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("bookmark with id: '%d' not found", id)
		}
		return "", err
	}

	// Check if the bookmark has a content ID
	if bookmark.ContentID != 0 {
		// Fetch the title from the student content table using the content ID
		var content models.Content
		err = repo.db.Where("id = ?", bookmark.ContentID).First(&content).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return "", fmt.Errorf("content with id: '%d' not found", bookmark.ContentID)
			}
			return "", err
		}
		return content.Title, nil
	}

	// Check if the bookmark has a question ID
	if bookmark.QuestionID != 0 {
		// Fetch the title from the question table using the question ID
		var question models.TypeQuestion
		err = repo.db.Where("id = ?", bookmark.QuestionID).First(&question).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return "", fmt.Errorf("question with id: '%d' not found", bookmark.QuestionID)
			}
			return "", err
		}
		return question.Title, nil
	}

	return "", fmt.Errorf("bookmark with id: '%d' does not have a content ID or question ID", id)
}

func (repo *Repository) GetBookmarkedContentAndCheckIfBookmarked(userID, contentID uint) (models.Bookmark, bool, error) {
	var bookmark models.Bookmark

	// Retrieve the bookmark from the database based on given id
	err := repo.db.Where("user_id = ? And content_id = ?", userID, contentID).First(&bookmark).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Bookmark{}, false, nil
		}

		return models.Bookmark{}, false, err
	}

	return bookmark, true, nil
}
func (repo *Repository) GetBookmarkedQuestionAndCheckIfBookmarked(userID, questionID uint) (models.Bookmark, bool, error) {
	var bookmark models.Bookmark

	// Retrieve the bookmark from the database based on given id
	err := repo.db.Where("user_id = ? And question_id = ?", userID, questionID).First(&bookmark).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Bookmark{}, false, nil
		}

		return models.Bookmark{}, false, err
	}

	return bookmark, true, nil
}
