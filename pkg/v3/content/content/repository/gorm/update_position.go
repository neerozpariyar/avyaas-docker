package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) UpdateContentPosition(data presenter.UpdateContentPositionRequest) map[string]string {
	errMap := make(map[string]string)
	transaction := repo.db.Begin()

	for idx, contentID := range data.ContentIDs {
		_, err := repo.GetContentByID(contentID)
		if err != nil {
			errMap["contentID"] = err.Error()
		}

		err = transaction.Model(&models.Content{}).Where("id = ?", contentID).Update("position", idx+1).Error
		if err != nil {
			errMap["contentID"] = err.Error()
		}
	}

	if len(errMap) != 0 {
		transaction.Rollback()
		return errMap
	}

	transaction.Commit()
	return errMap
}
