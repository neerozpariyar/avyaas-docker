package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"fmt"
)

func (repo *Repository) CreateNotice(data presenter.NoticeCreateUpdatePresenter) error {
	var err error

	notice := &models.Notice{
		Title:       data.Title,
		Description: data.Description,
		CreatedBy:   data.CreatedBy,
	}

	if data.File != nil {
		fileData, err := file.UploadFile("notice", data.File)
		if err != nil {
			return err
		}

		if fileData.FileType != "png" && fileData.FileType != "jpg" && fileData.FileType != "jpeg" && fileData.FileType != "pdf" { //validate file type before upload
			return fmt.Errorf("file type of %v not allowed:only image of type: png or jpg or jpeg or pdf", fileData.FileType)
		}

		isActive := true
		urlObject := utils.GetURLObject(fileData.Url)

		err = repo.db.Create(&models.File{
			Title:    fileData.Filename,
			Type:     fileData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			return err
		}

		notice.File = urlObject
	}

	if data.CourseID != 0 {
		notice.CourseID = data.CourseID
	}

	err = repo.db.Create(&notice).Error
	if err != nil {
		return err
	}

	return err

}
