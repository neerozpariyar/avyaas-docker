package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) GetTestHistory(request presenter.TestHistoryRequest) ([]models.TestResult, float64, error) {
	var results []models.TestResult
	var totalPage float64
	var err error
	var listCase string

	baseQuery := repo.db.Model(&models.TestResult{}).Where("user_id = ?", request.UserID)

	if request.CourseID != 0 {
		listCase = "course"
	} else if request.Type != "" {
		listCase = "type"
	} else if request.FromDate != "" || request.ToDate != "" {
		listCase = "dateRange"
	}

	switch listCase {
	case "course":
		println("course")
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = request.CourseID

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.TestResult{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if request.Page > int(totalPage) {
			request.Page = int(totalPage)
		}

		// Fetch a paginated list of tests from the database
		err = baseQuery.Where("course_id = ?", request.CourseID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&results).Error

		// return tests, totalPage, err
	case "type":
		println("type")
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["type"] = request.Type

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.TestResult{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if request.Page > int(totalPage) {
			request.Page = int(totalPage)
		}

		// Fetch a paginated list of tests from the database
		err = baseQuery.Where("type = ?", request.Type).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&results).Error

		// return tests, totalPage, err
	case "dateRange":
		println("dateRange")
		// Retrieve the tests based on the courseID provided
		if request.FromDate != "" && request.ToDate != "" {
			fd, _ := utils.ParseStringToTime(request.FromDate)
			ed, _ := utils.ParseStringToTime(request.ToDate)

			// Initialize an empty map to store condition data

			conditionData := make(map[string]interface{})
			conditionData["start_time"] = fd
			conditionData["end_time"] = ed

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Test{}, conditionData, true, []string{">=", "<="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err = baseQuery.Where("start_time >= ? AND end_time <= ?", fd, ed).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&results).Error

			// return tests, totalPage, err
		} else if request.FromDate != "" {
			fd, _ := utils.ParseStringToTime(request.FromDate)
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["start_time"] = fd

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.TestResult{}, conditionData, true, []string{">="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err = baseQuery.Where("start_time >= ?", fd).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&results).Error

			// return tests, totalPage, err
		} else if request.ToDate != "" {
			ed, _ := utils.ParseStringToTime(request.ToDate)
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["end_time"] = ed

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.TestResult{}, conditionData, true, []string{"<="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err = baseQuery.Where("end_time <= ?", ed).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&results).Error

			// return tests, totalPage, err
		}
	default:
		// Calculate the total number of pages based on the configured page size
		totalPage = utils.GetTotalPage(models.Test{}, repo.db, request.PageSize)

		if request.Page > int(totalPage) {
			request.Page = int(totalPage)
		}

		// Fetch a paginated list of tests from the database
		err = baseQuery.Preload("QuestionSets").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&results).Error

		// return tests, totalPage, err
	}

	return results, totalPage, err
}
