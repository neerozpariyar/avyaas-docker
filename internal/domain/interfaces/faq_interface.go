package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type FAQUsecase interface {
	CreateFaq(data *models.FAQ) error
	ListFaq(req presenter.FAQListReq) ([]models.FAQ, int64, error)
	UpdateFaq(data models.FAQ) (*models.FAQ, map[string]string)
	DeleteFaq(id uint) (*models.FAQ, error)
}

type FAQRepository interface {
	CreateFaq(data *models.FAQ) error
	GetFAQByID(id uint) (*models.FAQ, error)
	ListFaq(req presenter.FAQListReq) ([]models.FAQ, float64, error)
	UpdateFaq(data *models.FAQ) error
	DeleteFaq(id uint) (*models.FAQ, error)
}
