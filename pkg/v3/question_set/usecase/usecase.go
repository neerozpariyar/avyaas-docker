package usecase

import (
	"avyaas/internal/domain/interfaces"
)

/*
usecase represents the question set usecase, which contains the necessary components for handling
question set related business logic. It includes a question set repository for data access.
*/
type usecase struct {
	repo        interfaces.QuestionSetRepository
	accountRepo interfaces.AccountRepository
	courseRepo  interfaces.CourseRepository
	subjectRepo interfaces.SubjectRepository
	// questionRepo interfaces.QuestionRepository
	questionRepo interfaces.QuestionRepository
	bookmarkRepo interfaces.BookmarkRepository
}

/*
New initializes and returns a new instance of the question set usecase. It takes a question set
repository as parameter. The usecase is responsible for handling business logic related to question
set.
*/

func New(repo interfaces.QuestionSetRepository, accountRepo interfaces.AccountRepository, courseRepo interfaces.CourseRepository, subjectRepo interfaces.SubjectRepository, questionRepo interfaces.QuestionRepository, bookmarkRepo interfaces.BookmarkRepository) *usecase {
	return &usecase{
		repo:         repo,
		accountRepo:  accountRepo,
		courseRepo:   courseRepo,
		subjectRepo:  subjectRepo,
		questionRepo: questionRepo,
		bookmarkRepo: bookmarkRepo,
	}
}
