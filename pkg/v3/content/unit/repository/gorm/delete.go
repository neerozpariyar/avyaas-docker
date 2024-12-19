package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteUnit(id uint) error {
	transaction := repo.db.Begin()

	unit, err := repo.GetUnitByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if unit.Thumbnail != "" {
		var uFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", unit.Thumbnail).First(&uFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", uFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}
	// Perform a hard delete of the unit group with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&unit).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
