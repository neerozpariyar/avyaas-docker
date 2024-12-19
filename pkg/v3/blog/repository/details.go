package repository

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// func (repo *repository) DetailsOfBlog(id uint) (*models.Blog, error) {
// 	var blog *models.Blog
// 	err := repo.db.Where("id=?", id).First(&blog).Error

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, err
// 		}
// 		return nil, err
// 	}

// 	err = repo.db.Model(&models.Blog{}).Where("id = ?", id).Update("like", blog.Likes+1).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = repo.db.Model(&models.Blog{}).Where("id = ?", id).Update("views", blog.Views+1).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return blog, err
// }

func (repo *Repository) DetailsOfBlog(id uint) (*models.Blog, error) {
	var blog *models.Blog
	err := repo.db.Where("id=?", id).First(&blog).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("blog not found: %s", err)
		}
		return nil, fmt.Errorf("error fetching blog details: %s", err)
	}

	_, err = repo.GetCommentByBlogID(id)
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&models.Blog{}).Where("id = ?", id).Update("views", blog.Views+1).Error
	if err != nil {

		return nil, fmt.Errorf("error updating views: %s", err)
	}

	return blog, nil
}
