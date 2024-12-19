package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo interfaces.FAQRepository
}

func New(repo interfaces.FAQRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
