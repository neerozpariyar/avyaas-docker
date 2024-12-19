package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) CreateBlogComment(data models.BlogComment) (*models.Blog, error) {
	blogComment, err := uCase.repo.CreateBlogComment(data)
	if err != nil {
		return nil, err
	}

	return blogComment, err
}
