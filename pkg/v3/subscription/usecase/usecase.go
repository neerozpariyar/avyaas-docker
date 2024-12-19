package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo       interfaces.SubscriptionRepository
	courseRepo interfaces.CourseRepository
}

func New(repo interfaces.SubscriptionRepository, courseRepo interfaces.CourseRepository) *usecase {
	return &usecase{
		repo:       repo,
		courseRepo: courseRepo,
	}
}
