package usecase

import "avyaas/internal/domain/presenter"

func (uCase *usecase) UpdatePackageType(data presenter.PackageTypeCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing package  with the provided package's ID
	_, err = uCase.repo.GetPackageTypeByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	for _, serviceID := range data.ServiceIDs {
		_, err := uCase.serviceRepo.GetServiceByID(serviceID)
		if err != nil {
			errMap["serviceID"] = err.Error()
		}
	}

	if len(errMap) != 0 {
		return errMap
	}

	// Delegate the update of package type
	if err = uCase.repo.UpdatePackageType(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
