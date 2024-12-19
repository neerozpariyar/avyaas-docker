package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) DeleteBlog(id uint) (*models.Blog, error) {

	blog, err := uCase.repo.DeleteBlog(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}
