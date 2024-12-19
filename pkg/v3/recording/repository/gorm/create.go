package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"fmt"
)

func (repo *Repository) UploadRecording(data presenter.RecordingCreateUpdateRequest) error {
	recording := &models.Recording{
		Title:       data.Title,
		Description: data.Description,
		LiveID:      data.LiveID,
		Views:       data.Views,
		Length:      data.Length,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
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

		// rData, err := data.File.Open()
		// if err != nil {
		// 	return err
		// }

		// // Note: change the fle length check package later
		// pData, err := ffprobe.ProbeReader(context.Background(), rData)
		// if err != nil {
		// 	return err
		// }

		// recording.Length = uint(pData.Format.DurationSeconds)
		recording.Url = urlObject
	}

	var existingRecordings []models.Recording

	err := repo.db.Model(&models.Recording{}).Where("live_id = ?", data.LiveID).Find(&existingRecordings).Order("position").Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	var newRecording models.Recording
	err = transaction.Create(&recording).Scan(&newRecording).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	if len(existingRecordings) != 0 {
		err = transaction.Model(models.Recording{}).Where("id = ?", newRecording.ID).Update("position", existingRecordings[len(existingRecordings)-1].Position+1).Error

		if err != nil {
			transaction.Rollback()
			return err
		}
	} else {
		err = transaction.Model(models.Recording{}).Where("id = ?", newRecording.ID).Update("position", 1).Error

		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	transaction.Commit()
	return err
}
