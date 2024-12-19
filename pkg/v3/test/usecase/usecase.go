package usecase

import (
	"avyaas/internal/domain/interfaces"
	// "avyaas/pkg/v3/question_set"
)

/*
usecase represents the test usecase, which contains the necessary components for handling test
related business logic. It includes a test repository for data access.
*/
type usecase struct {
	repo            interfaces.TestRepository
	courseRepo      interfaces.CourseRepository
	questionSetRepo interfaces.QuestionSetRepository
	questionRepo    interfaces.QuestionRepository

	accountRepo    interfaces.AccountRepository
	testSeriesRepo interfaces.TestSeriesRepository

	questionSetUsecase interfaces.QuestionSetUsecase
}

/*
New initializes and returns a new instance of the test usecase. It takes a test repository as
parameter. The usecase is responsible for handling business logic related to test.
*/
func New(repo interfaces.TestRepository, courseRepo interfaces.CourseRepository, questionSetRepo interfaces.QuestionSetRepository, questionRepo interfaces.QuestionRepository, accountRepo interfaces.AccountRepository, testSeriesRepo interfaces.TestSeriesRepository, questionSetUsecase interfaces.QuestionSetUsecase) *usecase {
	return &usecase{
		repo:            repo,
		courseRepo:      courseRepo,
		questionSetRepo: questionSetRepo,
		questionRepo:    questionRepo,
		accountRepo:     accountRepo,
		testSeriesRepo:  testSeriesRepo,

		questionSetUsecase: questionSetUsecase,
	}
}
