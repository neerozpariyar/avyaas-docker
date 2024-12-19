package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (repo *Repository) UpdateBlog(requestBody presenter.BlogCreateUpdatePresenter) error {
	var err error
	updateBlog := &models.Blog{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		CreatedBy:   requestBody.CreatedBy,
		Tags:        requestBody.Tags,
	}

	if requestBody.Cover != nil {
		cover, err := file.UploadFile("notice", requestBody.Cover)
		if err != nil {
			return err
		}

		blog, err := repo.GetBlogByID(requestBody.ID)
		if err != nil {
			return err
		}

		isActive := true
		urlObject := utils.GetURLObject(cover.Url)

		err = repo.db.Create(&models.File{
			Title:    cover.Filename,
			Type:     cover.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error
		if err != nil {
			return err
		}

		if blog != nil {
			var data models.File

			err = repo.db.Debug().Model(&models.File{}).Where("url = ?", data.Url).Updates(&data).Error
			if err != nil {
				return err
			}
		}
	}

	if requestBody.CourseID != 0 {
		updateBlog.CourseID = requestBody.CourseID
	}
	if requestBody.SubjectID != 0 {
		updateBlog.SubjectID = requestBody.CourseID
	}

	err = repo.db.Where("id = ?", requestBody.ID).Updates(&updateBlog).Error
	if err != nil {
		return err
	}

	return err
}
