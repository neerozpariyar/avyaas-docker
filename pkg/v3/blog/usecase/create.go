package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error) {
	blog, err := uCase.repo.CreateBlog(data)
	if err != nil {
		return nil, err
	}

	return blog, err
}
