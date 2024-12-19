package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.SubjectRepository
	courseRepo  interfaces.CourseRepository
	accountRepo interfaces.AccountRepository
	contentRepo interfaces.ContentRepository
	unitRepo    interfaces.UnitRepository
}

func New(repo interfaces.SubjectRepository, courseRepo interfaces.CourseRepository, accountRepo interfaces.AccountRepository,
	contentRepo interfaces.ContentRepository, unitRepo interfaces.UnitRepository) *usecase {
	return &usecase{
		repo:        repo,
		courseRepo:  courseRepo,
		accountRepo: accountRepo,
		contentRepo: contentRepo,
		unitRepo:    unitRepo,
	}
}
