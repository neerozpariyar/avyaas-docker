package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListSubscriptions(request presenter.ListSubscriptionRequest) ([]models.Subscription, float64, error) {
	var subscriptions []models.Subscription
	baseQuery := repo.db.Debug().Model(&models.Subscription{}).Order("id")

	// Retrieve the subscriptions based on the courseID provided
	if request.CourseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionOperators := []string{}
		conditionData["course_id"] = request.CourseID
		conditionOperators = append(conditionOperators, "=")

		if request.Search != "" {
			conditionData["transaction_id"] = "%" + request.Search + "%"
			conditionOperators = append(conditionOperators, "like")
		}

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Subscription{}, conditionData, len(conditionOperators) == 1, conditionOperators, repo.db, request.PageSize)

		if request.Page > int(totalPage) {
			request.Page = int(totalPage)
		}

		// Fetch a paginated list of subscriptions from the database
		query := baseQuery
		for field, value := range conditionData {
			if field == "transaction_id" {
				query = query.Where(field+" like ?", value)
			} else {
				query = query.Where(field+" = ?", value)
			}
		}
		err := query.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&subscriptions).Error

		return subscriptions, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["transaction_id"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Subscription{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("transaction_id like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&subscriptions).Error

		return subscriptions, totalPage, err
	}

	if request.UserID != 0 {
		user, err := repo.accountRepo.GetUserByID(request.UserID)
		if err != nil {
			return subscriptions, 0, err
		}

		if user.RoleID == 4 {
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["user_id"] = request.UserID

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Subscription{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of subscriptions from the database
			err := baseQuery.Where("user_id = ?", request.UserID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&subscriptions).Error

			return subscriptions, totalPage, err
		}
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.Subscription{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	// Fetch a paginated list of subscriptions from the database
	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&subscriptions).Error

	return subscriptions, totalPage, err
}
