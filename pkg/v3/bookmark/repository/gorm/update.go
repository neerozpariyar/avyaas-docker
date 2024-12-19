package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"errors"

	"gorm.io/gorm"
)

func (repo *Repository) UpdateBookmark(data presenter.BookmarkCreateUpdateRequest) error {
	_, err := repo.GetBookmarkByID(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return repo.db.Where("id = ?", data.ID).Updates(&models.Bookmark{
		UserID:     data.UserID,
		QuestionID: data.QuestionID,
		ContentID:  data.ContentID,
	}).Error
}
