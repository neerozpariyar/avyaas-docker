package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) CreateUnit(data presenter.UnitCreateUpdateRequest) error {
	var urlObject string
	var currentUnit models.Unit

	unit := &models.Unit{
		Title:       data.Title,
		Description: data.Description,
		Thumbnail:   urlObject,
		// SubjectID:   data.SubjectID,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
		fileData, err := file.UploadFile("unit", data.File)
		if err != nil {
			return err
		}

		isActive := true
		urlObject = utils.GetURLObject(fileData.Url)

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

		unit.Thumbnail = urlObject
	}

	err := transaction.Create(&unit).Scan(&currentUnit).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	err = repo.subjectRepo.AssignUnitsToSubject(data.SubjectIDs, []uint{currentUnit.ID})

	transaction.Commit()
	return err
}
