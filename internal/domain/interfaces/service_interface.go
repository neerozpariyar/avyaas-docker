package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type ServiceUsecase interface {
	CreateService(data presenter.ServiceCreateUpdateRequest) map[string]string
	ListService(page int, search string, pageSize int) ([]models.Service, int, error)
	UpdateService(data presenter.ServiceCreateUpdateRequest) map[string]string
	DeleteService(id uint) error
}

type ServiceRepository interface {
	GetUrlByID(id uint) error

	CreateService(data presenter.ServiceCreateUpdateRequest) error
	ListService(page int, search string, pageSize int) ([]models.Service, float64, error)
	GetServiceByID(id uint) (*models.Service, error)
	GetServiceByTitle(title string) (*models.Service, error)
	GetUrlIDsByServiceID(serviceID uint) ([]uint, error)
	UpdateService(data presenter.ServiceCreateUpdateRequest) error
	DeleteService(id uint) error
}
