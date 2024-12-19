package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo interfaces.TermsAndConditionRepository
}

func New(repo interfaces.TermsAndConditionRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
