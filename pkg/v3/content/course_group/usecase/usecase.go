package usecase

import (
	"avyaas/internal/domain/interfaces"
)

/*
usecase represents the course group usecase, which contains the necessary components for handling
course group related business logic. It includes an course group repository for data access.
*/
type usecase struct {
	repo       interfaces.CourseGroupRepository
	courseRepo interfaces.CourseRepository
}

/*
New initializes and returns a new instance of the course group usecase. It takes a course group
repository as parameter. The usecase is responsible for handling business logic related to course
group.
*/
func New(repo interfaces.CourseGroupRepository, courseRepo interfaces.CourseRepository) *usecase {
	return &usecase{
		repo:       repo,
		courseRepo: courseRepo,
	}
}
