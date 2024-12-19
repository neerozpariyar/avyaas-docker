package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type ReferralUsecase interface {
	CreateReferral(data presenter.CreateUpdateReferralRequest) map[string]string
	ListReferral(request presenter.ReferralListRequest) ([]presenter.ReferralResponse, int, error)
	UpdateReferral(data presenter.CreateUpdateReferralRequest) map[string]string
	DeleteReferral(id uint) error

	ApplyReferral(request presenter.ApplyReferralRequest) (*presenter.ApplyReferralResponse, map[string]string)
}

type ReferralRepository interface {
	GetReferralByID(id uint) (models.Referral, error)
	GetReferralByCode(code string) (models.Referral, error)
	CheckPendingReferralInTransaction(id uint) (models.Referral, error)
	CreateReferral(data models.Referral) error
	ListReferral(request presenter.ReferralListRequest) ([]models.Referral, float64, error)
	UpdateReferral(data models.Referral) error
	DeleteReferral(id uint) error

	CheckUserReferral(userID, referralID uint) (*models.UserReferral, error)
	CreateReferralInTransaction(userID, referralID, packageID uint) error
}
