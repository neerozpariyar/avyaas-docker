package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) CreateUnit(data presenter.UnitCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	for _, subjectID := range data.SubjectIDs {
		if _, err := uCase.subjectRepo.GetSubjectByID(subjectID); err != nil {
			errMap["subjectID"] = err.Error()
			return errMap
		}
	}

	if err = uCase.repo.CreateUnit(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
