package usecase

import "avyaas/internal/domain/interfaces"

/*
usecase represents the account usecase, which contains the necessary components for handling account
related business logic. It includes an account repository for data access.
*/
type usecase struct {
	repo       interfaces.NoticeRepository
	courseRepo interfaces.CourseRepository
}

/*
New initializes and returns a new instance of the account usecase. It takes a account repository as
parameter. The usecase is responsible for handling business logic related to account.
*/
func New(repo interfaces.NoticeRepository, courseRepo interfaces.CourseRepository) *usecase {
	return &usecase{
		repo:       repo,
		courseRepo: courseRepo,
	}
}
