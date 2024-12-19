package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
)

/*
ListCourseGroup retrieves a paginated list of course groups from the database.

Parameters:
  - page: An integer representing the page number for pagination.

Returns:
  - data: A slice of Course Group representing the retrieved categories.
  - totalPage: A floating-point number representing the total number of pages available.
  - err: An error indicating the success or failure of the operation.
*/
// func (repo *repository) ListCourseGroup(page int, search string, pageSize int) ([]models.CourseGroup, float64, error) {
// 	var courseGroups []models.CourseGroup

// 	if search != "" {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["title"] = "%" + search + "%"
// 		conditionData["group_id"] = "%" + search + "%"

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.CourseGroup{}, conditionData, false, []string{"like", "like"}, repo.db, pageSize)

// 		if totalPage < float64(page) {
// 			page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.CourseGroup{}).Where("title like ? OR group_id like ?", "%"+search+"%", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Find(&courseGroups).Order("id").Error

// 		return courseGroups, totalPage, err
// 	}

// 	// Calculate the total number of pages based on the configured page size
// 	totalPage := utils.GetTotalPage(models.CourseGroup{}, repo.db, pageSize)

// 	if totalPage < float64(page) {
// 		page = int(totalPage)
// 	}

// 	// Fetch a paginated list of course groups from the database and preload associated courses
// 	err := repo.db.Debug().Model(&models.CourseGroup{}).Scopes(utils.Paginate(page, pageSize)).Preload("Courses").Find(&courseGroups).Order("id").Error

// 	return courseGroups, totalPage, err
// }

func (repo *Repository) ListCourseGroup(page int, search string, pageSize int) ([]models.CourseGroup, float64, error) {
	var courseGroups []models.CourseGroup

	baseQuery := repo.db.Debug().Model(&models.CourseGroup{}).Order("id")

	if search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + search + "%"
		conditionData["group_id"] = "%" + search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.CourseGroup{}, conditionData, false, []string{"like", "like"}, baseQuery, pageSize)

		if totalPage < float64(page) {
			page = int(totalPage)
		}

		err := baseQuery.Where("title like ? OR group_id like ?", "%"+search+"%", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Find(&courseGroups).Error

		return courseGroups, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.CourseGroup{}, baseQuery, pageSize)

	if totalPage < float64(page) {
		page = int(totalPage)
	}

	// Fetch a paginated list of course groups from the database and preload associated courses
	err := baseQuery.Scopes(utils.Paginate(page, pageSize)).Scopes(utils.Paginate(page, pageSize)).Preload("Courses").Find(&courseGroups).Error

	return courseGroups, totalPage, err
}
