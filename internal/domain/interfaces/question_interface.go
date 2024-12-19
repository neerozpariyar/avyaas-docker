package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

/*
QuestionUsecase represents the question usecase interface, defining methods for handling various question
related operations.
*/
type QuestionUsecase interface {
	CreateQuestion(question presenter.CreateUpdateQuestionRequest) map[string]string
	ListQuestion(request presenter.ListQuestionRequest) ([]presenter.QuestionListResponse, int, error)
	UpdateQuestion(data presenter.CreateUpdateQuestionRequest) map[string]string
	DeleteQuestion(id uint) error
}

/*
QuestionRepository represents the question repository interface, defining methods for handling various
question related operations.
*/
type QuestionRepository interface {
	GetQuestionByID(id uint) (models.Question, error)

	CreateQuestion(question presenter.CreateUpdateQuestionRequest) (*models.Question, error)
	ListQuestion(request presenter.ListQuestionRequest) ([]models.Question, float64, error)
	UpdateQuestion(data presenter.CreateUpdateQuestionRequest) error
	DeleteQuestion(id uint) error

	GetOptionByQuestionID(id uint) (*models.Option, error)
	GetOptionsByQuestionID(id uint) ([]*models.Option, error)
	CheckIsBookmarked(userID, contentID uint) (bool, error)
	GetCorrectAnswersIDForQuestion(queestionId uint) (uint, error)

	UpdateCaseQuestion(question presenter.CreateUpdateQuestionRequest) (*models.Question, error)
	UpdateFillInBlanksQuestion(question *presenter.CreateUpdateQuestionRequest) error
	UpdateMCQQuestion(question *presenter.CreateUpdateQuestionRequest) error
	UpdateTrueOrFalseQuestion(question *presenter.CreateUpdateQuestionRequest) error

	CreateCaseQuestion(question presenter.CreateUpdateQuestionRequest) (*models.Question, error)
	CreateFillInBlanksQuestion(question *presenter.CreateUpdateQuestionRequest) error
	CreateMCQQuestion(question *presenter.CreateUpdateQuestionRequest) error
	CreateTrueOrFalseQuestion(question *presenter.CreateUpdateQuestionRequest) error

	GetNestedQuestions(id uint) ([]presenter.QuestionListResponse, error)
}
