package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) SaveOrUpdateContentProgress(contentProgress *models.StudentContent) error {
	var existingContentProgress []models.StudentContent
	err := repo.db.Model(&models.StudentContent{}).Where("content_id = ?", contentProgress.ContentID).Find(&existingContentProgress).Error
	if err != nil {
		return err
	}

	if len(existingContentProgress) != 0 {
		// Update existing content progress
		return repo.db.Model(&models.StudentContent{}).
			Where("user_id = ? AND content_id = ?", contentProgress.UserID, contentProgress.ContentID).
			Updates(contentProgress).Error
	}

	// If doesn't exists create new content progress
	return repo.db.Create(contentProgress).Error
}
