package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) BlogLikeUnlike(uID, bID uint) (*models.Blog, error) {
	blog, err := uCase.repo.GetBlogByID(bID)
	if err != nil {
		return nil, err
	}

	user, err := uCase.accountRepo.GetUserByID(uID)
	if err != nil {
		return nil, err
	}

	blogLikeUnlike, err := uCase.repo.BlogLikeUnlike(user.ID, blog.ID)
	if err != nil {
		return nil, err
	}
	return blogLikeUnlike, err
}
