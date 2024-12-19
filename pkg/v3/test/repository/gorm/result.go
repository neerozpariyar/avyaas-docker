package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) GetTestResult(testID, requesterID uint) ([]models.TestResponse, error) {
	var responses []models.TestResponse

	err := repo.db.Where("user_id = ? AND test_id = ?", requesterID, testID).Find(&responses).Error
	if err != nil {
		return nil, err
	}

	return responses, nil
}
