package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) DeleteLive(id uint) error {
	transaction := repo.db.Begin()

	_, err := repo.GetLiveByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Live{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
