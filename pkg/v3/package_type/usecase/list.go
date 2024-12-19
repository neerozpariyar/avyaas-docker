package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListPackageType(request *presenter.PackageTypeListRequest) ([]models.PackageType, int, error) {
	packageTypes, totalPage, err := u.repo.ListPackageType(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	return packageTypes, int(totalPage), nil
}
