package usecase

import (
	"avyaas/internal/domain/interfaces"
)

type usecase struct {
	repo        interfaces.NotificationRepository
	courseRepo  interfaces.CourseRepository
	accountRepo interfaces.AccountRepository
}

func New(repo interfaces.NotificationRepository, courseRepo interfaces.CourseRepository, accountRepo interfaces.AccountRepository) *usecase {
	return &usecase{
		repo:        repo,
		courseRepo:  courseRepo,
		accountRepo: accountRepo,
	}
}
