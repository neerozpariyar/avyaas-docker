package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListTestSeries(request presenter.ListTestSeriesRequest) ([]models.TestSeries, float64, error) {
	var testSeries []models.TestSeries

	baseQuery := repo.db.Debug().Model(&models.TestSeries{}).Order("id")

	// Retrieve the tests based on the courseID provided
	if request.CourseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = request.CourseID

		if request.Search != "" {
			conditionData["title"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.TestSeries{}, conditionData, false, []string{"=", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&testSeries).Error

			return testSeries, totalPage, err
		}

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.TestSeries{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if request.Page > int(totalPage) {
			request.Page = int(totalPage)
		}

		// Fetch a paginated list of tests from the database
		err := baseQuery.Where("course_id = ?", request.CourseID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&testSeries).Error

		return testSeries, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.TestSeries{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&testSeries).Error

		return testSeries, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.TestSeries{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	// Fetch a paginated list of test series from the database
	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&testSeries).Error

	return testSeries, totalPage, err
}
