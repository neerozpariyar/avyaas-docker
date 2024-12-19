package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.ContentCommentRepository
	contentRepo interfaces.ContentRepository
	accountRepo interfaces.AccountRepository
}

func New(repo interfaces.ContentCommentRepository, contentRepo interfaces.ContentRepository, accountRepo interfaces.AccountRepository) *usecase {
	return &usecase{
		repo:        repo,
		contentRepo: contentRepo,
		accountRepo: accountRepo,
	}
}
