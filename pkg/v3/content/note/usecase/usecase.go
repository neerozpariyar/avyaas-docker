package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.NoteRepository
	contentRepo interfaces.ContentRepository
}

func New(repo interfaces.NoteRepository, contentRepo interfaces.ContentRepository) *usecase {
	return &usecase{
		repo:        repo,
		contentRepo: contentRepo,
	}
}
