package interfaces

// import (
// 	"avyaas/internal/domain/models"
// 	"avyaas/internal/domain/presenter"
// )

// /*
// QuestionUsecase represents the question usecase interface, defining methods for handling various question
// related operations.
// */
// type TypeQuestionUsecase interface {
// 	CreateTypeQuestion(question presenter.TypeQuestionPresenter) map[string]string
// 	UpdateTypeQuestion(data presenter.TypeQuestionPresenter) map[string]string
// 	DeleteTypeQuestion(id uint) error
// 	ListTypeQuestion(request presenter.ListQuestionRequest) ([]presenter.TypeQuestionListPresenter, int, error)
// 	// ListTypeQuestion(request presenter.ListQuestionRequest) ([]presenter.QuestionPresenter, int, error)
// 	// UpdateTypeQuestion(data presenter.CreateUpdateQuestionRequest) map[string]string
// 	// DeleteTypeQuestion(id uint) error
// 	GetQuestionTypeByID(id uint) (string, error)
// }

// /*
// QuestionRepository represents the question repository interface, defining methods for handling various
// question related operations.
// */
// type TypeQuestionRepository interface {
// 	// GetTypeQuestionByID(id uint) (models.Question, error)

// 	CreateCaseTypeQuestion(question presenter.TypeQuestionPresenter) (*models.TypeQuestion, error)
// 	CreateFillInBlanksQuestion(question *presenter.TypeQuestionPresenter) error
// 	CreateMCQQuestion(question *presenter.TypeQuestionPresenter) error
// 	CreateTrueOrFalseQuestion(question *presenter.TypeQuestionPresenter) error

// 	ListTypeQuestion(request presenter.ListQuestionRequest) ([]models.TypeQuestion, float64, error)

// 	DeleteTypeQuestion(id uint) error

// 	UpdateCaseTypeQuestion(question presenter.TypeQuestionPresenter) (*models.TypeQuestion, error)
// 	UpdateFillInBlanksQuestion(question *presenter.TypeQuestionPresenter) error
// 	UpdateMCQQuestion(question *presenter.TypeQuestionPresenter) error
// 	UpdateTrueOrFalseQuestion(question *presenter.TypeQuestionPresenter) error

// 	// ListTypeQuestion(request presenter.ListQuestionRequest) ([]models.Question, float64, error)
// 	// UpdateTypeQuestion(data presenter.CreateUpdateQuestionRequest) error
// 	// DeleteTypeQuestion(id uint) error

// 	GetTypeQuestionByID(id uint) (models.TypeQuestion, error)
// 	GetTypeOptionByQuestionID(id uint) (*models.TypeOption, error)
// 	GetTypeOptionsByQuestionID(questionID uint) ([]*models.TypeOption, error)
// 	CheckIsBookmarked(userID, questionID uint) (bool, error)
// 	GetNestedQuestions(id uint) ([]presenter.TypeQuestionListPresenter, error)
// }
