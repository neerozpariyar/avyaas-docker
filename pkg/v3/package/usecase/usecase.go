package usecase

import (
	"avyaas/internal/domain/interfaces"
)

/*
usecase represents the package usecase, which contains the necessary components for handling package
related business logic. It includes an package repository for data access.
*/
type usecase struct {
	repo            interfaces.PackageRepository
	packageTypeRepo interfaces.PackageTypeRepository
	courseRepo      interfaces.CourseRepository
	serviceRepo     interfaces.ServiceRepository
	testSeriesRepo  interfaces.TestSeriesRepository
	testRepo        interfaces.TestRepository
	liveGroupRepo   interfaces.LiveGroupRepository
	liveRepo        interfaces.LiveRepository
	paymentRepo     interfaces.PaymentRepository
}

/*
New initializes and returns a new instance of the package usecase. It takes a package repository as
parameter. The usecase is responsible for handling business logic related to package.
*/
func New(repo interfaces.PackageRepository, packageTypeRepo interfaces.PackageTypeRepository, courseRepo interfaces.CourseRepository, serviceRepo interfaces.ServiceRepository, testSeriesRepo interfaces.TestSeriesRepository, testRepo interfaces.TestRepository, liveGroupRepo interfaces.LiveGroupRepository, liveRepo interfaces.LiveRepository, paymentRepo interfaces.PaymentRepository) *usecase {
	return &usecase{
		repo:            repo,
		packageTypeRepo: packageTypeRepo,
		courseRepo:      courseRepo,
		serviceRepo:     serviceRepo,
		testSeriesRepo:  testSeriesRepo,
		liveGroupRepo:   liveGroupRepo,
		liveRepo:        liveRepo,
		paymentRepo:     paymentRepo,
	}
}
