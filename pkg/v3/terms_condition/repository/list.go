package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListTermsAndCondition(req presenter.TermsAndConditionListReq) ([]models.TermsAndCondition, float64, error) {
	var faq []models.TermsAndCondition
	var err error

	totalPage := utils.GetTotalPage(models.TermsAndCondition{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}
	err = repo.db.Debug().Model(models.TermsAndCondition{}).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&faq).Order("id").Error

	return faq, totalPage, err
}
