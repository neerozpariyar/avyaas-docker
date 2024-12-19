package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
)

/*
ListTestType retrieves a paginated list of test types from the database.

Parameters:
  - page: An integer representing the page number for pagination.

Returns:
  - data: A slice of models.TestType representing the retrieved test types.
  - totalPage: A floating-point number representing the total number of pages available.
  - err: An error indicating the success or failure of the operation.
*/
func (repo *Repository) ListTestType(page int, pageSize int) ([]models.TestType, float64, error) {
	var testTypes []models.TestType

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.TestType{}, repo.db, pageSize)

	if page > int(totalPage) {
		page = int(totalPage)
	}

	// Fetch a paginated list of test types from the database
	err := repo.db.Debug().Model(&models.TestType{}).Scopes(utils.Paginate(page, pageSize)).Find(&testTypes).Order("id").Error

	return testTypes, totalPage, err
}
