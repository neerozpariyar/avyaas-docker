package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) CreateSubject(data presenter.SubjectCreateUpdateRequest) error {
	subject := &models.Subject{
		SubjectID:   data.SubjectID,
		Title:       data.Title,
		Description: data.Description,
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

		subject.Thumbnail = urlObject
	}

	err := transaction.Create(&subject).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	if data.CourseIDs != nil {
		err = repo.courseRepo.AssignSubjectsToCourse(data.CourseIDs, []uint{subject.ID})
		if err != nil {
			return err
		}
	}

	return err
}
