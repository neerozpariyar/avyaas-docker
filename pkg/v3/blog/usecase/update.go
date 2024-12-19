package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) UpdateBlog(requestBody presenter.BlogCreateUpdatePresenter) (*models.Blog, map[string]string) {
	var err error

	errMap := make(map[string]string)

	blogID, err := uCase.repo.GetBlogByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateBlog(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return blogID, errMap
}
