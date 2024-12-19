package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) CreateNote(data presenter.NoteCreateUpdateRequest) error {
	note := &models.Note{
		Title:       data.Title,
		ContentID:   data.ContentID,
		Description: data.Description,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
		fileData, err := file.UploadFile("note", data.File)
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

		note.File = urlObject
	}

	if err := transaction.Create(&note).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
