package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"strings"
)

func (repo *Repository) UpdateContent(data presenter.ContentCreateUpdateRequest) error {
	updatedContent := &models.Content{
		Title:       data.Title,
		Description: data.Description,
		IsPremium:   data.IsPremium,
		ContentType: strings.ToUpper(data.ContentType),
		Length:      data.Length,
		Level:       data.Level,
		Visibility:  data.Visibility,
		CreatedBy:   data.CreatedBy,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
		fileData, err := file.UploadFile("content", data.File)
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

		content, err := repo.GetContentByID(data.ID)
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

		// if strings.ToUpper(data.ContentType) == "VIDEO" {
		// 	// Note: change the fle length check package later
		// 	pData, err := ffprobe.ProbeReader(context.Background(), rData)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	updatedContent.Length = uint(pData.Format.DurationSeconds)
		// }

		updatedContent.Url = urlObject
	}

	err := repo.db.Where("id = ?", data.ID).Updates(&updatedContent).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
