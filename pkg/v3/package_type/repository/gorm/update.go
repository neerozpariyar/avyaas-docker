package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
)

func (repo *Repository) UpdatePackageType(data presenter.PackageTypeCreateUpdateRequest) error {
	transaction := repo.db.Begin()

	err := transaction.Where("id = ?", data.ID).Updates(&models.PackageType{
		Title:       data.Title,
		Description: data.Description,
	}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	var oldServiceIDs []uint
	err = repo.db.Select("service_id").Table("package_type_services").Where("package_type_id = ?", data.ID).Scan(&oldServiceIDs).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	addIDs, delIDs := utils.CompareDifferences(oldServiceIDs, data.ServiceIDs)

	if len(addIDs) > 0 {
		for _, id := range addIDs {
			query := fmt.Sprintf(`INSERT INTO package_type_services (package_type_id, service_id) VALUES (%d, %d);`, data.ID, id)
			err = transaction.Exec(query).Error
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	if len(delIDs) > 0 {
		for _, id := range delIDs {
			query := fmt.Sprintf(`DELETE FROM package_type_services WHERE package_type_id = %d AND service_id = %d;`, data.ID, id)
			err = transaction.Debug().Exec(query).Error
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	transaction.Commit()
	return err
}
