package repository

import "avyaas/internal/domain/models"

func (repo *Repository) CreateBlogComment(data models.BlogComment) (*models.Blog, error) {
	comment := &models.BlogComment{
		Comment: data.Comment,
	}

	if data.BlogID != 0 {
		comment.BlogID = data.BlogID
	}
	if data.UserID != 0 {
		comment.UserID = data.UserID
	}

	err := repo.db.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return nil, err
}
