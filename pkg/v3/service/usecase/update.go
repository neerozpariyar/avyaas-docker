package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) UpdateService(data presenter.ServiceCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing service with the provided service's ID
	service, err := uCase.repo.GetServiceByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Check if a service  with the given title already exists
	serviceByTitle, err := uCase.repo.GetServiceByTitle(data.Title)
	if err == nil {
		// Check if the title is the same as of the requested service
		if service.Title != serviceByTitle.Title {
			errMap["title"] = fmt.Errorf("service with title: '%s' already exists", serviceByTitle.Title).Error()
			return errMap
		}
	}

	// Delegate the update of service
	if err = uCase.repo.UpdateService(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
