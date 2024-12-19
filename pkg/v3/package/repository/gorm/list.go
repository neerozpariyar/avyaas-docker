package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListPackage(request *presenter.PackageListRequest) ([]models.Package, float64, error) {
	var packages []models.Package
	baseQuery := repo.db.Debug().Model(&models.Package{}).Order("id")

	if request.CourseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = request.CourseID
		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["title"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Package{}, conditionData, false, []string{"=", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("title LIKE ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&packages).Error

			return packages, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Package{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("course_id = ?", request.CourseID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&packages).Error

		return packages, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Package{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("title LIKE ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&packages).Error

		return packages, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.Package{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&packages).Error

	return packages, totalPage, err

}
