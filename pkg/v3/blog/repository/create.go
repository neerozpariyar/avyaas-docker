package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"fmt"
)

func (repo *Repository) CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error) {
	var err error

	blog := &models.Blog{
		Title:       data.Title,
		Description: data.Description,
		Tags:        data.Tags,
		CreatedBy:   data.CreatedBy,
	}

	if data.Cover != nil {
		coverData, err := file.UploadFile("blog", data.Cover)
		if err != nil {
			return nil, err
		}

		isActive := true
		urlObject := utils.GetURLObject(coverData.Url)

		err = repo.db.Create(&models.File{
			Title:    coverData.Filename,
			Type:     coverData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			return nil, err
		}
		fmt.Printf("urlObject: %v\n", urlObject)
		blog.Cover = urlObject
	}

	if data.CourseID != 0 {
		blog.CourseID = data.CourseID
	}

	if data.SubjectID != 0 {
		blog.SubjectID = data.SubjectID
	}

	err = repo.db.Create(&blog).Error
	if err != nil {
		return nil, err
	}

	return nil, err
}
