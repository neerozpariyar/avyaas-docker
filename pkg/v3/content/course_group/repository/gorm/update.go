package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

/*
UpdateCourse is a repository method responsible for updating the details of a course group in the
repository based on the provided course group's ID.

Parameters:
  - courseGroup: A models.CourseGroup instance containing the updated details of the course group.

Returns:
  - err: An error, if any, encountered during the update operation.
*/
func (repo *Repository) UpdateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) error {
	updateCourse := &models.CourseGroup{
		GroupID:     data.GroupID,
		Title:       data.Title,
		Description: data.Description,
	}

	transaction := repo.db.Begin()

	if data.File != nil {

		fileData, err := file.UploadFile("course_group", data.File)
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

		courseGroup, err := repo.GetCourseGroupByID(data.ID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if courseGroup.Thumbnail != "" {
			var cgFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", courseGroup.Thumbnail).First(&cgFile).Error
			if err == nil { // If there was an existing thumbnail, delete it from storage
				if err = repo.db.Model(models.File{}).Where("id = ?", cgFile.ID).Update("is_active", false).Error; err != nil {
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
