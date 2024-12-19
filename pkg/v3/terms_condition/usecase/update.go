package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdateTermsAndCondition(data models.TermsAndCondition) (*models.TermsAndCondition, map[string]string) {
	var err error
	errMap := make(map[string]string)

	termID, err := uCase.repo.GetTermsAndConditionByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateTermsAndCondition(&data)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return termID, errMap
}
