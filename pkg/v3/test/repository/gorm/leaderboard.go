package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) GetTestLeaderboard(testID uint) ([]models.TestResult, error) {
	var results []models.TestResult
	if err := repo.db.Where("test_id = ?", testID).Order("score desc").Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
