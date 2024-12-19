package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteBookmark(id uint) error {
	transaction := repo.db.Begin()

	// Perform a hard delete of the bookmark group with the given ID using the GORM Unscoped method
	err := transaction.Unscoped().Where("id = ?", id).Delete(&models.Bookmark{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
