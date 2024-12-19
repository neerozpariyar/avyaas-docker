package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteService(id uint) error {
	transaction := repo.db.Begin()

	// err := transaction.Table("package_services").Unscoped().Where("service_id = ?", id).Delete(&models.Service{}).Error
	err := transaction.Exec("DELETE FROM package_services WHERE service_id = ?", id).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the service with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Service{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	err = transaction.Where("service_id = ?", id).Delete(&models.ServiceUrl{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
