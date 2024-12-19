package usecase

import (
	"avyaas/internal/domain/interfaces"
)

/*
usecase represents the question usecase, which contains the necessary components for handling
question related business logic. It includes a question repository for data access.
*/
type usecase struct {
	repo            interfaces.TypeQuestionRepository
	courseRepo      interfaces.CourseRepository
	bookmarkRepo    interfaces.BookmarkRepository
	subjectRepo     interfaces.SubjectRepository
	questionSetRepo interfaces.QuestionSetRepository
}

/*
New initializes and returns a new instance of the question usecase. It takes a question repository
as parameter. The usecase is responsible for handling business logic related to question.
*/
func New(repo interfaces.TypeQuestionRepository, courseRepo interfaces.CourseRepository, subjectRepo interfaces.SubjectRepository, questionSetRepo interfaces.QuestionSetRepository, bookmarkRepo interfaces.BookmarkRepository) *usecase {
	return &usecase{
		repo:            repo,
		courseRepo:      courseRepo,
		subjectRepo:     subjectRepo,
		questionSetRepo: questionSetRepo,
		bookmarkRepo:    bookmarkRepo,
	}
}
