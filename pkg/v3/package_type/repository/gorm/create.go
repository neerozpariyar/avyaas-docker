package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (repo *Repository) CreatePackageType(data presenter.PackageTypeCreateUpdateRequest) error {
	var packageType *models.PackageType
	transaction := repo.db.Begin()

	err := transaction.Create(&models.PackageType{
		Title:       data.Title,
		Description: data.Description,
	}).Scan(&packageType).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	for _, serviceID := range data.ServiceIDs {
		query := fmt.Sprintf(`INSERT INTO package_type_services (package_type_id, service_id) VALUES (%d, %d);`, packageType.ID, serviceID)
		err = transaction.Exec(query).Error
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	transaction.Commit()
	return err
}
