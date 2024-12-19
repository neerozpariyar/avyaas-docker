package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"math"
)

/*
ListTest retrieves a paginated list of test from the database.

Parameters:
  - page: An integer representing the page number for pagination.
  - courseID: A uint representing the courseID to filter the test.

Returns:
  - data: A slice of models.Test representing the retrieved tests.
  - totalPage: A floating-point number representing the total number of pages available.
  - err: An error indicating the success or failure of the operation.
*/
func (repo *Repository) ListTest(request presenter.ListTestRequest) ([]models.Test, float64, error) {
	var tests []models.Test
	var totalPage float64
	baseQuery := repo.db.Model(&models.Test{}).Order("start_time")

	var listCase string

	if request.Status != "" {
		listCase = "status"
	} else if request.CourseID != 0 {
		listCase = "course"
	} else if request.TestTypeID != 0 {
		listCase = "type"
	} else if request.FromDate != "" || request.ToDate != "" {
		listCase = "dateRange"
	}

	user, err := repo.accountRepo.GetUserByID(request.UserID)
	if err != nil {
		return tests, totalPage, err
	}

	if user.RoleID == 4 {
		if request.TestTypeID != 0 {
			conditionData := make(map[string]interface{})
			conditionData["user_id"] = request.UserID
			conditionData["course_id"] = request.CourseID
			conditionData["test_type_id"] = request.TestTypeID

			err = baseQuery.Debug().Where("id IN (?) AND test_type_id = ?", repo.db.Select("test_id").Model(&models.StudentTest{}).Where("user_id = ? AND course_id = ?", request.UserID, request.CourseID), request.TestTypeID).Find(&tests).Error
			if err != nil {
				return tests, totalPage, err
			}

			var freeTests []models.Test
			err = baseQuery.Where("is_free = ? AND test_type_id = ? AND course_id = ?", true, request.TestTypeID, request.CourseID).Find(&freeTests).Error
			if err != nil {
				return nil, totalPage, err
			}

			tests = append(tests, freeTests...)

			totalPage = math.Ceil(float64(len(tests)) / float64(request.PageSize))
			return tests, totalPage, err
		}

		conditionData := make(map[string]interface{})
		conditionData["user_id"] = request.UserID
		conditionData["course_id"] = request.CourseID

		// Calculate the total number of pages based on the configured page size
		// totalPage := utils.GetTotalPageByConditionModel(models.StudentTest{}, conditionData, true, []string{"=", "="}, repo.db, request.PageSize)

		// if request.Page > int(totalPage) {
		// 	request.Page = int(totalPage)
		// }

		err = baseQuery.Where("id IN (?)", repo.db.Select("test_id").Model(&models.StudentTest{}).Where("user_id = ? AND course_id = ?", request.UserID, request.CourseID)).Find(&tests).Error
		if err != nil {
			return nil, totalPage, err
		}

		var freeTests []models.Test
		err = baseQuery.Where("is_free = ? AND course_id = ?", true, request.CourseID).Find(&freeTests).Error
		if err != nil {
			return nil, totalPage, err
		}

		tests = append(tests, freeTests...)

		totalPage = math.Ceil(float64(len(tests)) / float64(request.PageSize))

		return tests, totalPage, err
	}

	switch listCase {
	case "course":
		println("course")
		// Retrieve the tests based on the courseID provided
		if request.CourseID != 0 {
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["course_id"] = request.CourseID

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Test{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err = baseQuery.Where("course_id = ?", request.CourseID).Preload("QuestionSets").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&tests).Error

			// return tests, totalPage, err
		}
	case "status":
		println("status")
		// Retrieve the tests based on the test status
		if request.Status != "" {
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["status"] = request.Status

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Test{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err = baseQuery.Where("status = ?", request.Status).Preload("QuestionSets").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&tests).Error

			// return tests, totalPage, err
		}
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
			err = baseQuery.Where("start_time >= ?", fd).Preload("QuestionSets").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&tests).Error

			// return tests, totalPage, err
		} else if request.FromDate != "" {
			fd, _ := utils.ParseStringToTime(request.FromDate)
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["start_time"] = fd

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Test{}, conditionData, true, []string{">="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err = baseQuery.Where("start_time >= ?", fd).Preload("QuestionSets").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&tests).Error

			// return tests, totalPage, err
		} else if request.ToDate != "" {
			ed, _ := utils.ParseStringToTime(request.ToDate)
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["end_time"] = ed

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Test{}, conditionData, true, []string{"<="}, repo.db, request.PageSize)

			if request.Page > int(totalPage) {
				request.Page = int(totalPage)
			}

			// Fetch a paginated list of tests from the database
			err = baseQuery.Where("end_time <= ?", ed).Preload("QuestionSets").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&tests).Error

			// return tests, totalPage, err
		}
	default:
		// Calculate the total number of pages based on the configured page size
		totalPage = utils.GetTotalPage(models.Test{}, repo.db, request.PageSize)

		if request.Page > int(totalPage) {
			request.Page = int(totalPage)
		}

		// Fetch a paginated list of tests from the database
		err = baseQuery.Preload("QuestionSets").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&tests).Error

		// return tests, totalPage, err
	}

	return tests, totalPage, err
}
