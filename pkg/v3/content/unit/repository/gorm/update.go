package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) UpdateUnit(data presenter.UnitCreateUpdateRequest) error {
	updatedUnit := &models.Unit{
		Title: data.Title,
		// Description: data.Description,
		// SubjectID:   data.SubjectID,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
		fileData, err := file.UploadFile("unit", data.File)
		if err != nil {
			return err
		}

		isActive := true
		urlObject := utils.GetURLObject(fileData.Url)
		err = transaction.Create(&models.File{
			Title:    fileData.Filename,
			Type:     fileData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		unit, err := repo.GetUnitByID(data.ID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if unit.Thumbnail != "" {
			var uFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", unit.Thumbnail).First(&uFile).Error
			if err == nil {
				if err = repo.db.Model(models.File{}).Where("id = ?", uFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		updatedUnit.Thumbnail = urlObject
	}

	err := transaction.Where("id = ?", data.ID).Updates(&updatedUnit).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
