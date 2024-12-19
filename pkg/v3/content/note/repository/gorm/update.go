package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) UpdateNote(data presenter.NoteCreateUpdateRequest) error {
	updatedNote := models.Note{
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
			Url:      fileData.Url,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		note, err := repo.GetNoteByID(data.ID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if note.File != "" {
			var nFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", note.File).First(&nFile).Error
			if err == nil {
				if err = repo.db.Model(models.File{}).Where("id = ?", nFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		updatedNote.File = urlObject
	}

	err := repo.db.Where("id = ?", data.ID).Updates(&updatedNote).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
