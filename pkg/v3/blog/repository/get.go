package repository

import "avyaas/internal/domain/models"

func (repo *Repository) GetBlogByID(id uint) (*models.Blog, error) {
	var blog models.Blog

	err := repo.db.Where("id = ?", id).First(&blog).Error
	if err != nil {
		return nil, err
	}

	return &blog, err
}

func (repo *Repository) GetCommentByID(id uint) (*models.BlogComment, error) {
	var comment models.BlogComment

	err := repo.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, err
}

func (repo *Repository) GetCommentByBlogID(id uint) (*models.BlogComment, error) {
	var comment models.BlogComment

	err := repo.db.Where("blog_id", comment.BlogID).Find(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, err
}
