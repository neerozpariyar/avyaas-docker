package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

/*
Usecase represents the question set usecase interface, defining methods for handling various question
set related operations.
*/
type QuestionSetUsecase interface {
	CreateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) map[string]string
	ListQuestionSet(page int, courseID uint, pageSize int) ([]presenter.QuestionSetDetailsPresenter, int, error)
	GetQuestionSetDetails(id, requesterID uint) (*presenter.QuestionSetDetailsPresenter, error)
	UpdateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) map[string]string
	AssignQuestionsToQuestionSet(questionSetID uint, questionIDs []uint) error
	DeleteQuestionSet(id uint) error
}

/*
Repository represents the question set repository interface, defining methods for handling various
question set related operations.
*/
type QuestionSetRepository interface {
	GetQuestionSetByID(id uint) (models.QuestionSet, error)
	GetQuestionSetByTitleAndCourseID(title string, courseID uint) (models.QuestionSet, error)
	CreateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) error
	ListQuestionSet(page int, courseID uint, pageSize int) ([]models.QuestionSet, float64, error)
	UpdateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) error
	DeleteQuestionSet(id uint) error
	AssignQuestionsToQuestionSet(questionSetID uint, questionIDs []uint) error
	GetQuestionSetQuestion(questionSetID, questionID uint) (*models.QuestionSetQuestion, error)
}
