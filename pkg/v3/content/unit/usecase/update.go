package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) UpdateUnit(data presenter.UnitCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	for _, subjectID := range data.SubjectIDs {
		if _, err := uCase.subjectRepo.GetSubjectByID(subjectID); err != nil {
			errMap["subjectID"] = err.Error()
			return errMap
		}
	}

	// Retrieve the existing unit  with the provided unit 's ID
	_, err = uCase.repo.GetUnitByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Delegate the update of unit
	if err = uCase.repo.UpdateUnit(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
