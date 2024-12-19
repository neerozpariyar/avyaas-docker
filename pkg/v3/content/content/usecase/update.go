package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) UpdateContent(data presenter.ContentCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing content  with the provided content 's ID
	_, err = uCase.repo.GetContentByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Delegate the update of content
	if err = uCase.repo.UpdateContent(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
