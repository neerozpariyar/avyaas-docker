package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) ListFaq(req presenter.FAQListReq) ([]models.FAQ, int64, error) {
	faqs, totalPage, err := uCase.repo.ListFaq(req)
	if err != nil {
		return nil, int64(totalPage), err
	}
	var allFaq []models.FAQ

	for _, faq := range faqs {
		allFaq = append(allFaq, models.FAQ{
			Timestamp: models.Timestamp{
				ID: faq.ID,
			},
			Title:       faq.Title,
			Description: faq.Description,
		})
	}
	return allFaq, int64(totalPage), err
}
