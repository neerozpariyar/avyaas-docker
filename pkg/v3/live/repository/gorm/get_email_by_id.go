package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) GetEmailByID(id uint) (string, error) {
	var live models.Live
	if err := repo.db.First(&live, id).Error; err != nil {
		return "", fmt.Errorf("failed to fetch email: %v", err)
	}
	return live.Email, nil
}
