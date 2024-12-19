package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo           interfaces.ReplyRepository
	courseRepo     interfaces.CourseRepository
	discussionRepo interfaces.DiscussionRepository
	accountRepo    interfaces.AccountRepository
}

func New(repo interfaces.ReplyRepository, courseRepo interfaces.CourseRepository, discussionRepo interfaces.DiscussionRepository, accountRepo interfaces.AccountRepository) *usecase {
	return &usecase{
		repo:           repo,
		courseRepo:     courseRepo,
		discussionRepo: discussionRepo,
		accountRepo:    accountRepo,
	}
}
