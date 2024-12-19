package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo            interfaces.LiveGroupRepository
	courseRepo      interfaces.CourseRepository
	packageTypeRepo interfaces.PackageTypeRepository
}

func New(repo interfaces.LiveGroupRepository, courseRepo interfaces.CourseRepository, packageTypeRepo interfaces.PackageTypeRepository) *usecase {
	return &usecase{
		repo:            repo,
		courseRepo:      courseRepo,
		packageTypeRepo: packageTypeRepo,
	}
}
