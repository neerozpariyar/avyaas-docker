package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListComments(requestBody presenter.BlogCommentListReq) ([]models.BlogComment, float64, error) {
	var comments []models.BlogComment

	baseQuery := repo.db.Debug().Model(&models.BlogComment{})

	if requestBody.UserID != 0 {
		baseQuery = baseQuery.Where("course_id = ?", requestBody.UserID)
		if requestBody.Search != "" {
			baseQuery = baseQuery.Where("title like ? OR user_id like ?", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%")
		}
	}

	if requestBody.BlogID != 0 {
		baseQuery = baseQuery.Where("subject_id = ?", requestBody.BlogID)
		if requestBody.Search != "" {
			baseQuery = baseQuery.Where("title like ? OR blog_id like ?", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%")
		}
	}

	if requestBody.Search != "" {
		baseQuery = baseQuery.Where("title like ? OR user_id like ? OR blog_id like?", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%")
	}

	totalPage := utils.GetTotalPage(models.BlogComment{}, baseQuery, requestBody.PageSize)

	if totalPage < float64(requestBody.Page) {
		requestBody.Page = int(totalPage)
	}
	err := baseQuery.Scopes(utils.Paginate(requestBody.Page, requestBody.PageSize)).Find(&comments).Order("id").Error

	return comments, totalPage, err
}
