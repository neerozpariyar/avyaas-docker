package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) UpdateSubject(data presenter.SubjectCreateUpdateRequest) error {
	updatedSubject := &models.Subject{
		SubjectID:   data.SubjectID,
		Title:       data.Title,
		Description: data.Description,
		// CourseID:    data.CourseID,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
		fileData, err := file.UploadFile("subject", data.File)
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

		subject, err := repo.GetSubjectByID(data.ID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if subject.Thumbnail != "" {
			var sFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", subject.Thumbnail).First(&sFile).Error
			if err == nil {
				if err = repo.db.Model(models.File{}).Where("id = ?", sFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		updatedSubject.Thumbnail = urlObject
	}

	return repo.db.Where("id = ?", data.ID).Updates(&updatedSubject).Error
}
