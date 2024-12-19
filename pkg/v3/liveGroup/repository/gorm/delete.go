package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteLiveGroup(id uint) error {
	transaction := repo.db.Begin()

	_, err := repo.GetLiveGroupByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the liveGroup  with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.LiveGroup{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
