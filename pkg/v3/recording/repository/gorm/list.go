package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

// func (repo *repository) ListRecording(request presenter.RecordingListRequest) ([]models.Recording, float64, error) {
// 	var recordings []models.Recording
// 	if request.LiveID != 0 {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["live_id"] = request.LiveID

// 		if request.Search != "" {
// 			// Initialize an empty map to store condition data
// 			conditionData["title"] = "%" + request.Search + "%"
// 			// Calculate the total number of pages based on the configured page size for given filter condition
// 			totalPage := utils.GetTotalPageByConditionModel(models.Recording{}, conditionData, false, []string{"=", "like"}, repo.db, request.PageSize)

// 			if totalPage < float64(request.Page) {
// 				request.Page = int(totalPage)
// 			}

// 			err := repo.db.Debug().Model(&models.Recording{}).Where("live_id=? AND title like ?", request.LiveID, "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&recordings).Error

// 			return recordings, totalPage, err
// 		}
// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Recording{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

// 		if totalPage < float64(request.Page) {
// 			request.Page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Recording{}).Where("live_id = ?", request.LiveID).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("position").Find(&recordings).Error

// 		return recordings, totalPage, err
// 	}

// 	if request.Search != "" {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["title"] = "%" + request.Search + "%"

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Recording{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

// 		if totalPage < float64(request.Page) {
// 			request.Page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Recording{}).Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&recordings).Error

// 		return recordings, totalPage, err
// 	}

// 	totalPage := utils.GetTotalPage(models.Recording{}, repo.db, request.PageSize)

// 	if totalPage < float64(request.Page) {
// 		request.Page = int(totalPage)
// 	}

// 	err := repo.db.Debug().Model(&models.Recording{}).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&recordings).Error

// 	return recordings, totalPage, err

// }

func (repo *Repository) ListRecording(request presenter.RecordingListRequest) ([]models.Recording, float64, error) {
	var recordings []models.Recording

	// Initialize base query
	baseQuery := repo.db.Debug().Model(&models.Recording{})

	// Initialize an empty map to store condition data
	conditionData := make(map[string]interface{})
	conditionOperators := []string{"="}

	if request.Search != "" {
		conditionData["title"] = "%" + request.Search + "%"
		conditionOperators = []string{"like"}
	}
	if request.LiveID != 0 {

		conditionData["live_id"] = request.LiveID
		conditionOperators = append(conditionOperators, "=")
		if request.Search != "" {
			conditionData["title"] = "%" + request.Search + "%"
			conditionOperators = append(conditionOperators, "like")
		}
	}

	//Calculate the total number of pages based on the configured page size for given filter condition
	totalPage := utils.GetTotalPageByConditionModel(models.Recording{}, conditionData, len(conditionOperators) > 1, conditionOperators, repo.db, request.PageSize)
	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	query := baseQuery
	// Add the conditions to the query
	for field, value := range conditionData {
		if field == "title" {
			query = query.Where(field+" like ?", value)
		} else {
			query = query.Where(field+" = ?", value)
		}
	}

	// Order the query
	if request.LiveID != 0 || request.Search != "" {
		query = query.Order("id")
	} else {
		query = query.Order("position")
	}

	err := query.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&recordings).Error
	return recordings, totalPage, err
}
