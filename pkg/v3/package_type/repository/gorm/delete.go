package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) DeletePackageType(id uint) error {
	transaction := repo.db.Begin()

	query := fmt.Sprintf(`DELETE FROM package_type_services WHERE package_type_id = %d;`, id)
	err := transaction.Exec(query).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the package type with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.PackageType{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
