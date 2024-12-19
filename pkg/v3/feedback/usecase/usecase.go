package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo       interfaces.FeedbackRepository
	courseRepo interfaces.CourseRepository
}

func New(repo interfaces.FeedbackRepository, courseRepo interfaces.CourseRepository) *usecase {
	return &usecase{
		repo:       repo,
		courseRepo: courseRepo,
	}
}
