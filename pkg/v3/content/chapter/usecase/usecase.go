package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.ChapterRepository
	unitRepo    interfaces.UnitRepository
	contentRepo interfaces.ContentRepository
}

func New(repo interfaces.ChapterRepository, unitRepo interfaces.UnitRepository, contentRepo interfaces.ContentRepository) *usecase {
	return &usecase{
		repo:        repo,
		unitRepo:    unitRepo,
		contentRepo: contentRepo,
	}
}
