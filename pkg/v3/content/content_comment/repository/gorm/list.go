package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListComment(request presenter.CommentListRequest) ([]models.Comment, float64, error) {
	var comments []models.Comment
	baseQuery := repo.db.Debug().Model(&models.Comment{}).Order("id")

	if request.ContentID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["content_id"] = request.ContentID

		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["comment"] = "%" + request.Search + "%"
			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Comment{}, conditionData, true, []string{"=", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("content_id = ? AND comment like ?", request.ContentID, "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&comments).Error
			return comments, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Comment{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("content_id = ?", request.ContentID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&comments).Error

		return comments, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = request.Search

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Comment{}, conditionData, true, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("comment like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&comments).Error

		return comments, totalPage, err
	}

	totalPage := utils.GetTotalPage(models.Comment{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&comments).Error

	return comments, totalPage, err

}
