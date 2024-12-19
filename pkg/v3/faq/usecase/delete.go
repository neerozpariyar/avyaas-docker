package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) DeleteFaq(id uint) (*models.FAQ, error) {
	faq, err := uCase.repo.DeleteFaq(id)
	if err != nil {
		return nil, err
	}
	return faq, nil
}
