package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetPackageTypeByID(id uint) (models.PackageType, error) {
	var packageType models.PackageType

	// Retrieve the package type from the database based on given id
	err := repo.db.Where("id = ?", id).First(&packageType).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.PackageType{}, fmt.Errorf("package type with id: '%d' not found", id)
		}

		return models.PackageType{}, err
	}

	return packageType, nil
}

func (repo *Repository) GetPackageTypeServices(packageTypeID uint) ([]uint, error) {
	var serviceIDs []uint

	err := repo.db.Table("package_type_services").Select("service_id").Where("package_type_id = ?", packageTypeID).Scan(&serviceIDs).Error
	return serviceIDs, err
}
