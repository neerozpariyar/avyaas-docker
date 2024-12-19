package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.BookmarkRepository
	contentRepo interfaces.ContentRepository
	// questionRepo interfaces.QuestionRepository
	questionRepo interfaces.QuestionRepository
}

func New(repo interfaces.BookmarkRepository, contentRepo interfaces.ContentRepository, questionRepo interfaces.QuestionRepository) *usecase {
	return &usecase{
		repo:         repo,
		contentRepo:  contentRepo,
		questionRepo: questionRepo,
	}
}
