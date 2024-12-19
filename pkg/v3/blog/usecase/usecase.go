package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.BlogRepository
	accountRepo interfaces.AccountRepository
	courseRepo  interfaces.CourseRepository
	subjectRepo interfaces.SubjectRepository
}

func New(repo interfaces.BlogRepository, accountRepo interfaces.AccountRepository, courseRepo interfaces.CourseRepository, subjectRepo interfaces.SubjectRepository) *usecase {
	return &usecase{
		repo:        repo,
		accountRepo: accountRepo,
		courseRepo:  courseRepo,
		subjectRepo: subjectRepo,
	}
}
