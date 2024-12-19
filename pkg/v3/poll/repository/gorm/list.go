package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListPoll(request presenter.PollListRequest) ([]models.Poll, float64, error) {
	var polls []models.Poll
	//var replies []models.Reply
	baseQuery := repo.db.Debug().Model(&models.Poll{}).Order("id")

	if request.SubjectID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["subject_id"] = request.SubjectID
		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["question"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Poll{}, conditionData, false, []string{"=", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("question like ?", "%"+request.Search+"%").Preload("Options").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&polls).Error

			return polls, totalPage, err
		}

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Poll{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("subject_id = ?", request.SubjectID).Preload("Options").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&polls).Error

		return polls, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["question"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Poll{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("question like ?", "%"+request.Search+"%").Preload("Options").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&polls).Error

		return polls, totalPage, err
	}

	totalPage := utils.GetTotalPage(models.Poll{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	// err := repo.db.Debug().Model(&models.Poll{}).Preload("Options", models.PollOption{}).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&polls).Error
	err := baseQuery.Preload("Options").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&polls).Error
	return polls, totalPage, err

}
