package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListNotice(req presenter.NoticeListReq) ([]models.Notice, float64, error) {

	var notices []models.Notice

	baseQuery := repo.db.Debug().Model(&models.Notice{})

	if req.CourseID != 0 {
		baseQuery = baseQuery.Where("course_id = ?", req.CourseID)
		if req.Search != "" {
			baseQuery = baseQuery.Where("title like ? OR course_id like ?", "%"+req.Search+"%", "%"+req.Search+"%")
		}
	}

	if req.Search != "" {
		baseQuery = baseQuery.Where("title like ? OR course_id like ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	totalPage := utils.GetTotalPage(models.Notice{}, baseQuery, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}
	err := baseQuery.Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&notices).Order("id").Error

	return notices, totalPage, err

}
