package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) DeletePackage(id uint) error {
	transaction := repo.db.Begin()

	query := fmt.Sprintf(`DELETE FROM package_services WHERE package_id = %d;`, id)
	err := transaction.Exec(query).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the package with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Package{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
