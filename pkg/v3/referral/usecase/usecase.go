package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.ReferralRepository
	accountRepo interfaces.AccountRepository
	courseRepo  interfaces.CourseRepository
	packageRepo interfaces.PackageRepository
}

func New(repo interfaces.ReferralRepository, accountRepo interfaces.AccountRepository, courseRepo interfaces.CourseRepository, packageRepo interfaces.PackageRepository) *usecase {
	return &usecase{
		repo:        repo,
		accountRepo: accountRepo,
		courseRepo:  courseRepo,
		packageRepo: packageRepo,
	}
}
