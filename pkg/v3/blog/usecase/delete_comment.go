package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) DeleteBlogComment(id uint) (*models.BlogComment, error) {

	comment, err := uCase.repo.DeleteBlogComment(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
