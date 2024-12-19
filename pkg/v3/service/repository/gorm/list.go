package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
)

func (repo *Repository) ListService(page int, search string, pageSize int) ([]models.Service, float64, error) {
	var services []models.Service
	baseQuery := repo.db.Debug().Model(&models.Service{}).Order("id")

	if search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Service{}, conditionData, false, []string{"like"}, repo.db, pageSize)

		if totalPage < float64(page) {
			page = int(totalPage)
		}

		err := baseQuery.Where("title LIKE ?", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Find(&services).Order("id").Error

		return services, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.Service{}, repo.db, pageSize)

	if totalPage < float64(page) {
		page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(page, pageSize)).Find(&services).Order("id").Error

	return services, totalPage, err

}
