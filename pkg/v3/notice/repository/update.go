package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) UpdateNotice(requestBody presenter.NoticeCreateUpdatePresenter) error {
	var err error

	updateNotice := &models.Notice{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		CreatedBy:   requestBody.CreatedBy,
	}

	if requestBody.File != nil {
		file, err := file.UploadFile("notice", requestBody.File)
		if err != nil {
			return err
		}

		notice, err := repo.GetNoticeByID(requestBody.ID)
		if err != nil {
			return err
		}

		isActive := true
		urlObject := utils.GetURLObject(file.Url)

		err = repo.db.Create(&models.File{
			Title:    file.Filename,
			Type:     file.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error
		if err != nil {
			return err
		}

		if notice != nil {
			var data models.File

			err = repo.db.Debug().Model(&models.File{}).Where("url = ?", data.Url).Updates(&data).Error
			if err != nil {
				return err
			}
		}
	}

	if requestBody.CourseID != 0 {
		updateNotice.CourseID = requestBody.CourseID
	}

	err = repo.db.Where("id = ?", requestBody.ID).Updates(&updateNotice).Error
	if err != nil {
		return err
	}

	return err

}
