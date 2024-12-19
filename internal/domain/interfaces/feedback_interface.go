package interfaces

import "avyaas/internal/domain/models"

type FeedbackUsecase interface {
	CreateFeedback(data models.Feedback) map[string]string
	ListFeedback(page int, courseGroupID uint, pageSize int) ([]models.Feedback, int, error)
	UpdateFeedback(feedback models.Feedback) map[string]string
	DeleteFeedback(id uint) error
}

type FeedbackRepository interface {
	GetFeedbackByID(id uint) (models.Feedback, error)
	CreateFeedback(data models.Feedback) error
	ListFeedback(page int, courseGroupID uint, pageSize int) ([]models.Feedback, float64, error)
	UpdateFeedback(feedback models.Feedback) error
	DeleteFeedback(id uint) error
}
