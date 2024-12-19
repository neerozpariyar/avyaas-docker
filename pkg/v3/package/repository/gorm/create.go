package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

/*
CreatePackage is a repository function that persists the provided package in the database.

Parameters:
  - test: A pointer to the models.Package structure containing information about the package to be
    created.

Returns:
  - error: An error, if any, encountered during the database operation. Returns nil on success.
*/
func (repo *Repository) CreatePackage(data presenter.PackageCreateUpdateRequest) error {
	// var newPackage *models.Package
	// transaction := repo.db.Begin()

	err := repo.db.Create(&models.Package{
		Title:         data.Title,
		Description:   data.Description,
		PackageTypeID: data.PackageTypeID,
		CourseID:      data.CourseID,
		TestSeriesID:  data.TestSeriesID,
		TestID:        data.TestID,
		LiveGroupID:   data.LiveGroupID,
		LiveID:        data.LiveID,
		Price:         data.Price,
		Period:        data.Period,
		// Discount:        data.Discount,
		// DiscountedPrice: data.DiscountedPrice,
		// }).Scan(&newPackage).Error
	}).Error

	// if err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }

	// for _, serviceID := range data.ServiceIDs {
	// 	query := fmt.Sprintf(`INSERT INTO package_services (package_id, service_id) VALUES (%d, %d);`, newPackage.ID, serviceID)
	// 	err = transaction.Exec(query).Error
	// 	if err != nil {
	// 		transaction.Rollback()
	// 		return err
	// 	}
	// }

	// transaction.Commit()
	return err
}
