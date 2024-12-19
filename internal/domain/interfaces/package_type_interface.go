package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type PackageTypeUsecase interface {
	CreatePackageType(data presenter.PackageTypeCreateUpdateRequest) map[string]string
	ListPackageType(request *presenter.PackageTypeListRequest) ([]models.PackageType, int, error)
	UpdatePackageType(data presenter.PackageTypeCreateUpdateRequest) map[string]string
	DeletePackageType(id uint) error

	GetPackageTypeServices(packageTypeID uint) ([]uint, error)
}

type PackageTypeRepository interface {
	GetPackageTypeByID(id uint) (models.PackageType, error)
	CreatePackageType(data presenter.PackageTypeCreateUpdateRequest) error
	ListPackageType(request *presenter.PackageTypeListRequest) ([]models.PackageType, float64, error)
	UpdatePackageType(data presenter.PackageTypeCreateUpdateRequest) error
	DeletePackageType(id uint) error

	GetPackageTypeServices(packageTypeID uint) ([]uint, error)
}
