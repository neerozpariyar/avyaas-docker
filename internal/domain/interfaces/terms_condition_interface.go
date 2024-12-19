package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type TermsAndConditionUsecase interface {
	CreateTermsAndCondition(data *models.TermsAndCondition) error
	ListTermsAndCondition(req presenter.TermsAndConditionListReq) ([]models.TermsAndCondition, int64, error)
	UpdateTermsAndCondition(data models.TermsAndCondition) (*models.TermsAndCondition, map[string]string)
	DeleteTermsAndCondition(id uint) (*models.TermsAndCondition, error)
}

type TermsAndConditionRepository interface {
	CreateTermsAndCondition(data *models.TermsAndCondition) error
	GetTermsAndConditionByID(id uint) (*models.TermsAndCondition, error)
	ListTermsAndCondition(req presenter.TermsAndConditionListReq) ([]models.TermsAndCondition, float64, error)
	UpdateTermsAndCondition(data *models.TermsAndCondition) error
	DeleteTermsAndCondition(id uint) (*models.TermsAndCondition, error)
}
