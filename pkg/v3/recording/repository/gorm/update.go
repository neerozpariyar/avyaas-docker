package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"

	"fmt"
)

func (repo *Repository) UpdateRecording(data presenter.RecordingCreateUpdateRequest) error {
	if data.File != nil {
		transaction := repo.db.Begin()

		fileData, err := file.UploadFile("recording", data.File)
		if err != nil {
			return err
		}

		if fileData.FileType != "video/mp4" { //validate file type before upload
			return fmt.Errorf("file type of %v not allowed: only VIDEO of type: mp4 allowed", fileData.FileType)
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

		content, err := repo.GetRecordingByID(data.ID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if content.Url != "" {
			var cFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", content.Url).First(&cFile).Error
			if err == nil {
				if err = repo.db.Model(models.File{}).Where("id = ?", cFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		// rData, err := data.File.Open()
		// if err != nil {
		// 	return err
		// }

		// // Note: change the fle length check package later
		// pData, err := ffprobe.ProbeReader(context.Background(), rData)
		// if err != nil {
		// 	return err
		// }

		err = transaction.Updates(&models.Recording{
			Title:       data.Title,
			Description: data.Description,
			LiveID:      data.LiveID,
			Length:      data.Length,
			Views:       data.Views,
			Url:         urlObject,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		transaction.Commit()
		return err
	}

	return repo.db.Where("id = ?", data.ID).Updates(&models.Recording{
		LiveID:      data.LiveID,
		Title:       data.Title,
		Description: data.Description,
		Views:       data.Views,
		Length:      data.Length,
	}).Error
}
