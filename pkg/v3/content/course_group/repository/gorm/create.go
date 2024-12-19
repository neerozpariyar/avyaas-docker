package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

/*
CreateCourseGroup is a repository method responsible for creating a new course group in the database.

Parameters:
  - courseGroup: A models.CourseGroup instance representing the course group to be created in the database.

Returns:
  - error: An error, if any, encountered during the database insertion operation.
*/
func (repo *Repository) CreateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) error {
	courseGroup := &models.CourseGroup{
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

		courseGroup.Thumbnail = urlObject
	}

	err := transaction.Create(&courseGroup).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
