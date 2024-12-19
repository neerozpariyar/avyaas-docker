package usecase

import (
	"avyaas/internal/domain/interfaces"
)

type usecase struct {
	repo            interfaces.CourseRepository
	courseGroupRepo interfaces.CourseGroupRepository
	packageRepo     interfaces.PackageRepository
	paymentRepo     interfaces.PaymentRepository
	accountRepo     interfaces.AccountRepository
	subjectRepo     interfaces.SubjectRepository
}

func New(repo interfaces.CourseRepository, courseGroupRepo interfaces.CourseGroupRepository, packageRepo interfaces.PackageRepository, paymentRepo interfaces.PaymentRepository, accountRepo interfaces.AccountRepository, subjectRepo interfaces.SubjectRepository) *usecase {
	return &usecase{
		repo:            repo,
		courseGroupRepo: courseGroupRepo,
		packageRepo:     packageRepo,
		paymentRepo:     paymentRepo,
		accountRepo:     accountRepo,
		subjectRepo:     subjectRepo,
	}
}
