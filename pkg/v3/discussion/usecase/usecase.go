package usecase

import (
	"avyaas/internal/domain/interfaces"
)

type usecase struct {
	repo        interfaces.DiscussionRepository
	courseRepo  interfaces.CourseRepository
	subjectRepo interfaces.SubjectRepository
	accountRepo interfaces.AccountRepository
}

func New(repo interfaces.DiscussionRepository, courseRepo interfaces.CourseRepository, subjectRepo interfaces.SubjectRepository, accountRepo interfaces.AccountRepository) *usecase {
	return &usecase{
		repo:        repo,
		courseRepo:  courseRepo,
		subjectRepo: subjectRepo,
		accountRepo: accountRepo,
	}
}
