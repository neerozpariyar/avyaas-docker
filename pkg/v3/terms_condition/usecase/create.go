package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) CreateTermsAndCondition(data *models.TermsAndCondition) error {
	var err error

	termsAndCondition := uCase.repo.CreateTermsAndCondition(data)
	if err != nil {
		return err
	}
	return termsAndCondition
}
