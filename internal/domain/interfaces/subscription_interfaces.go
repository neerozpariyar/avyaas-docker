package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type SubscriptionUsecase interface {
	ListSubscriptions(request presenter.ListSubscriptionRequest) ([]presenter.SubscriptionListResponse, int, error)
}

type SubscriptionRepository interface {
	ListSubscriptions(request presenter.ListSubscriptionRequest) ([]models.Subscription, float64, error)
}
