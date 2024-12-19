package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetPackageByID(id uint) (models.Package, error) {
	var packageData models.Package

	// Retrieve the package from the database based on given id
	err := repo.db.Where("id = ?", id).First(&packageData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Package{}, fmt.Errorf("package with package id: '%d' not found", id)
		}

		return models.Package{}, err
	}

	return packageData, nil
}

func (repo *Repository) GetPackageByTestSeriesID(id uint) (models.Package, error) {
	var packageData models.Package

	// Retrieve the package from the database based on given id
	err := repo.db.Where("test_series_id = ? AND test_id = ? AND live_group_id = ? AND live_id = ?", id, 0, 0, 0).First(&packageData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Package{}, fmt.Errorf("package with package id: '%d' not found", id)
		}

		return models.Package{}, err
	}

	return packageData, nil
}

func (repo *Repository) CheckCoursePackage(courseID, packageID uint) (models.Package, error) {
	var packageData models.Package

	// Retrieve the package from the database associated to given course id
	err := repo.db.Where("id = ? AND course_id = ?", packageID, courseID).First(&packageData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Package{}, fmt.Errorf("package '%d' is not associated with course '%d'", packageID, courseID)
		}

		return models.Package{}, err
	}

	return packageData, nil
}

func (repo *Repository) GetCourseIDByPackageID(packageID uint) (uint, error) {
	var courseID uint
	// Retrieve the course ID from the package based on given package ID
	err := repo.db.Model(&models.Package{}).Select("course_id").Where("id = ?", packageID).Scan(&courseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("package with package id: '%d' not found", packageID)
		}
		return 0, err
	}
	return courseID, nil
}
