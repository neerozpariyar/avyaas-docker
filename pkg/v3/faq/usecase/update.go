package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdateFaq(data models.FAQ) (*models.FAQ, map[string]string) {
	var err error

	errMap := make(map[string]string)

	faqID, err := uCase.repo.GetFAQByID(data.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateFaq(&data)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}

	return faqID, errMap

}
