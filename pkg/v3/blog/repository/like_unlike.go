package repository

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) BlogLikeUnlike(uID, bID uint) (*models.Blog, error) {
	var likes models.BlogLike
	var blog models.Blog

	if err := repo.db.Where("user_id = ? AND blog_id = ?", uID, bID).First(&likes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			like := models.BlogLike{
				UserID: uID,
				BlogID: bID,
			}

			if err := repo.db.Create(&like).Error; err != nil {
				return nil, err
			}

			err = repo.db.Where("id = ?", bID).First(&blog).Update("likes", blog.Likes+1).Error
			if err != nil {
				return nil, err

			}
		}
		return nil, err

	}
	if &likes != nil {
		err := repo.db.Where("blog_id = ?", bID).Delete(&likes).Error
		if err != nil {
			return nil, fmt.Errorf("error updating likes: %s", err)
		}
		err = repo.db.Where("id = ?", bID).First(&blog).Update("likes", blog.Likes-1).Error
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &blog, nil
}
