package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListReply(request presenter.ReplyListRequest) ([]models.Reply, float64, error) {
	var replies []models.Reply
	baseQuery := repo.db.Debug().Model(&models.Reply{}).Order("id")

	if request.DiscussionID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["discussion_id"] = request.DiscussionID

		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["reply"] = "%" + request.Search + "%"
			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Reply{}, conditionData, true, []string{"=", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("discussion_id = ? AND reply like ?", request.DiscussionID, "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&replies).Error
			return replies, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Reply{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("discussion_id = ?", request.DiscussionID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&replies).Error

		return replies, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = request.Search

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Reply{}, conditionData, true, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("reply like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&replies).Error

		return replies, totalPage, err
	}

	totalPage := utils.GetTotalPage(models.Reply{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&replies).Error

	return replies, totalPage, err

}
