package usecase

import (
	"avyaas/internal/domain/interfaces"
)

/*
usecase represents the payment usecase, which contains the necessary components for handling payment
related business logic. It includes an payment repository for data access.
*/
type usecase struct {
	repo           interfaces.PaymentRepository
	accountRepo    interfaces.AccountRepository
	courseRepo     interfaces.CourseRepository
	packageUsecase interfaces.PackageUsecase
	packageRepo    interfaces.PackageRepository
}

/*
New initializes and returns a new instance of the payment usecase. It takes a payment repository as
parameter. The usecase is responsible for handling business logic related to payment.
*/
func New(repo interfaces.PaymentRepository, accountRepo interfaces.AccountRepository, courseRepo interfaces.CourseRepository, packageUsecase interfaces.PackageUsecase, packageRepo interfaces.PackageRepository) *usecase {
	return &usecase{
		repo:           repo,
		accountRepo:    accountRepo,
		courseRepo:     courseRepo,
		packageUsecase: packageUsecase,
		packageRepo:    packageRepo,
	}
}
