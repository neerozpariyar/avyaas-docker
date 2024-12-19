package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) UpdatePackage(data presenter.PackageCreateUpdateRequest) error {
	err := repo.db.Where("id = ?", data.ID).Updates(&models.Package{
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
	}).Error
	if err != nil {
		return err
	}

	// var oldServiceIDs []uint
	// err = repo.db.Select("service_id").Table("package_services").Where("package_id = ?", data.ID).Scan(&oldServiceIDs).Error
	// if err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }

	// addIDs, delIDs := utils.CompareDifferences(oldServiceIDs, data.ServiceIDs)s

	// if len(addIDs) > 0 {
	// 	for _, id := range addIDs {
	// 		query := fmt.Sprintf(`INSERT INTO package_services (package_id, service_id) VALUES (%d, %d);`, data.ID, id)
	// 		err = transaction.Exec(query).Error
	// 		if err != nil {
	// 			transaction.Rollback()
	// 			return err
	// 		}
	// 	}
	// }

	// if len(delIDs) > 0 {
	// 	for _, id := range delIDs {
	// 		query := fmt.Sprintf(`DELETE FROM package_services WHERE package_id = %d AND service_id = %d;`, data.ID, id)
	// 		err = transaction.Debug().Exec(query).Error
	// 		if err != nil {
	// 			transaction.Rollback()
	// 			return err
	// 		}
	// 	}
	// }

	return err
}
