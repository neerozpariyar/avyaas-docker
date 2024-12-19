package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) CreatePackageType(data presenter.PackageTypeCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	for _, serviceID := range data.ServiceIDs {
		_, err := uCase.serviceRepo.GetServiceByID(serviceID)
		if err != nil {
			errMap["serviceID"] = err.Error()
		}
	}

	if len(errMap) != 0 {
		return errMap
	}

	// Call the repository to create the package
	if err = uCase.repo.CreatePackageType(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
