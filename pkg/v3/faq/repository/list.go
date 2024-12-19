package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListFaq(req presenter.FAQListReq) ([]models.FAQ, float64, error) {
	var faq []models.FAQ
	var err error

	totalPage := utils.GetTotalPage(models.FAQ{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}
	err = repo.db.Debug().Model(models.FAQ{}).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&faq).Order("id").Error

	return faq, totalPage, err
}
