package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) UpdateCourse(data presenter.CourseCreateUpdateRequest) error {
	updateCourse := &models.Course{
		CourseID:    data.CourseID,
		Title:       data.Title,
		Description: data.Description,
		Available:   &data.Available,
		// CourseGroupID: data.CourseGroupID,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
		fileData, err := file.UploadFile("course", data.File)
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

		course, err := repo.GetCourseByID(data.ID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if course.Thumbnail != "" {
			var cFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", course.Thumbnail).First(&cFile).Error
			if err == nil { // If there was an existing thumbnail, delete it from storage
				if err = repo.db.Model(models.File{}).Where("id = ?", cFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		updateCourse.Thumbnail = urlObject
	}

	err := transaction.Where("id = ?", data.ID).Updates(&updateCourse).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
