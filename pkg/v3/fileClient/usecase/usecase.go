package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo interfaces.FileClientRepository
}

func New(repo interfaces.FileClientRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
