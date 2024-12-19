package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo            interfaces.TestSeriesRepository
	courseRepo      interfaces.CourseRepository
	packageTypeRepo interfaces.PackageTypeRepository
	packageRepo     interfaces.PackageRepository
}

func New(repo interfaces.TestSeriesRepository, courseRepo interfaces.CourseRepository, packageTypeRepo interfaces.PackageTypeRepository, packageRepo interfaces.PackageRepository) *usecase {
	return &usecase{
		repo:            repo,
		courseRepo:      courseRepo,
		packageTypeRepo: packageTypeRepo,
		packageRepo:     packageRepo,
	}
}
