package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) UpdateTestStatus(test models.Test) error {
	var status string

	if *test.IsPublic {
		status = "Inactive"
	} else {
		status = "Scheduled"
	}

	if err := repo.db.Debug().Model(&models.Test{}).Where("id = ?", test.ID).Updates(&models.Test{
		IsPublic: test.IsPublic,
		Status:   status,
	}).Error; err != nil {
		return err
	}

	return nil
}
