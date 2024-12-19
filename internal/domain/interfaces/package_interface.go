package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type PackageUsecase interface {
	CreatePackage(data presenter.PackageCreateUpdateRequest) map[string]string
	ListPackage(request *presenter.PackageListRequest) ([]presenter.PackageListResponse, int, error)
	UpdatePackage(data presenter.PackageCreateUpdateRequest) map[string]string
	DeletePackage(id uint) error

	SubscribePackage(request presenter.SubscribePackageRequest) map[string]string
}

type PackageRepository interface {
	CreatePackage(data presenter.PackageCreateUpdateRequest) error
	GetPackageByID(id uint) (models.Package, error)
	GetPackageByTestSeriesID(id uint) (models.Package, error)
	GetCourseIDByPackageID(packageID uint) (uint, error)
	CheckCoursePackage(courseID, packageID uint) (models.Package, error)
	ListPackage(request *presenter.PackageListRequest) ([]models.Package, float64, error)
	UpdatePackage(data presenter.PackageCreateUpdateRequest) error
	DeletePackage(id uint) error

	SubscribePackage(request presenter.SubscribePackageRequest) error
}
