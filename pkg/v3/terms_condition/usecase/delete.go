package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) DeleteTermsAndCondition(id uint) (*models.TermsAndCondition, error) {
	termsAndCondition, err := uCase.repo.DeleteTermsAndCondition(id)
	if err != nil {
		return nil, err
	}
	return termsAndCondition, err
}
