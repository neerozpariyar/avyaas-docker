package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) ListTermsAndCondition(req presenter.TermsAndConditionListReq) ([]models.TermsAndCondition, int64, error) {
	terms, totalPage, err := uCase.repo.ListTermsAndCondition(req)
	if err != nil {
		return nil, int64(totalPage), err
	}
	var allTerms []models.TermsAndCondition

	for _, term := range terms {
		allTerms = append(allTerms, models.TermsAndCondition{
			Timestamp: models.Timestamp{
				ID: term.ID,
			},
			Title:       term.Title,
			Description: term.Description,
		})
	}
	return allTerms, int64(totalPage), err
}
