package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListPackageType(request *presenter.PackageTypeListRequest) ([]models.PackageType, float64, error) {
	var packageTypes []models.PackageType
	baseQuery := repo.db.Debug().Model(&models.PackageType{}).Order("id")

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.PackageType{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("title LIKE ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Preload("Services").Find(&packageTypes).Error

		return packageTypes, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.PackageType{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Preload("Services").Find(&packageTypes).Error

	return packageTypes, totalPage, err

}
