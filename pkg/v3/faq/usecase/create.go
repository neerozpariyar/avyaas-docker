package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) CreateFaq(data *models.FAQ) error {
	var err error

	faq := uCase.repo.CreateFaq(data)
	if err != nil {
		return nil
	}
	return faq
}
