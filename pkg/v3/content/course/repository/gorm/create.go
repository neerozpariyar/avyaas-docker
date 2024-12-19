package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) CreateCourse(data presenter.CourseCreateUpdateRequest) error {
	course := &models.Course{
		CourseID:    data.CourseID,
		Title:       data.Title,
		Description: data.Description,
		Available:   &data.Available,
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

		course.Thumbnail = urlObject
	}

	if err := transaction.Create(&course).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	if data.CourseGroupIDs != nil {

		err := repo.courseGroupRepo.AssignCoursesToCourseGroup(data.CourseGroupIDs, []uint{course.ID})
		if err != nil {
			return err
		}

	}

	return nil
}
