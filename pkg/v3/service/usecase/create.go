package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) CreateService(data presenter.ServiceCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// for _, urlID := range data.UrlIDs {
	// 	if err = uCase.repo.GetUrlByID(urlID); err != nil {
	// 		uID := utils.UintToString(urlID)
	// 		errMap[uID] = err.Error()
	// 	}
	// }

	// if len(errMap) != 0 {
	// 	return errMap
	// }

	// Call the repository to create a new service
	if err = uCase.repo.CreateService(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
