package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListPayments(request *presenter.PaymentListRequest) ([]models.Payment, float64, error) {
	var payments []models.Payment

	user, err := repo.accountRepo.GetUserByID(request.UserID)
	if err != nil {
		return payments, 0, err
	}

	baseQuery := repo.db.Debug().Model(&models.Payment{}).Order("created_at")

	if user.RoleID == 4 { // means the user is student
		// baseQuery := repo.db.Debug().Model(&models.Payment{}).Where("user_id = ?", request.UserID).Order("created_at")

		// Retrieve the tests based on the courseID provided
		if request.CourseID != 0 {
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["course_id"] = request.CourseID
			conditionData["user_id"] = request.UserID

			if request.Search != "" {
				// Initialize an empty map to store condition data
				conditionData["transaction_id"] = "%" + request.Search + "%"

				// Calculate the total number of pages based on the configured page size for given filter condition
				totalPage := utils.GetTotalPageByConditionModel(models.Payment{}, conditionData, false, []string{"=", "=", "like"}, repo.db, request.PageSize)

				if totalPage < float64(request.Page) {
					request.Page = int(totalPage)
				}

				err := baseQuery.Where("transaction_id LIKE ? AND user_id = ?", "%"+request.Search+"%", request.UserID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

				return payments, totalPage, err
			}
			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Payment{}, conditionData, true, []string{"=", "="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err := baseQuery.Where("course_id = ? AND user_id = ?", request.CourseID, request.UserID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

			return payments, totalPage, err
		}

		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["transaction_id"] = "%" + request.Search + "%"
			conditionData["user_id"] = request.UserID

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Payment{}, conditionData, false, []string{"like", "="}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("transaction_id LIKE ? AND user_id = ?", "%"+request.Search+"%", request.UserID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

			return payments, totalPage, err
		}

		conditionData := make(map[string]interface{})
		conditionData["user_id"] = request.UserID

		// Calculate the total number of pages based on the configured page size
		// totalPage := utils.GetTotalPage(models.Payment{}, repo.db, request.PageSize)
		totalPage := utils.GetTotalPageByConditionModel(models.Payment{}, conditionData, false, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err = baseQuery.Where("user_id = ?", request.UserID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

		return payments, totalPage, err
	}

	// baseQuery := repo.db.Debug().Model(&models.Payment{}).Order("created_at")

	// Retrieve the tests based on the courseID provided
	if request.CourseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = request.CourseID
		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["transaction_id"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Payment{}, conditionData, false, []string{"=", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("transaction_id LIKE ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

			return payments, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Payment{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if request.Page > int(totalPage) {
			request.Page = int(totalPage)
		}

		// Fetch a paginated list of tests from the database
		err := baseQuery.Where("course_id = ?", request.CourseID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

		return payments, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["transaction_id"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Payment{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("transaction_id LIKE ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

		return payments, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.Payment{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	err = baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&payments).Error

	return payments, totalPage, err

}
