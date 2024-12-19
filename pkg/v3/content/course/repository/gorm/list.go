package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListCourse(request presenter.CourseListRequest) ([]models.Course, float64, error) {
	var courses []models.Course

	user, err := repo.accountRepo.GetUserByID(request.UserID)
	if err != nil {
		return courses, 0, err
	}

	if request.CourseGroupID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_group_id"] = request.CourseGroupID
		if request.Search != "" {

			conditionData["title"] = "%" + request.Search + "%"
			conditionData["course_id"] = "%" + request.Search + "%"
			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, false, []string{"=", "like", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := repo.db.Debug().Model(&models.Course{}).Where("title like ? OR course_id like ?", "%"+request.Search+"%", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

			return courses, totalPage, err
		}
		if user.RoleID == 4 {
			conditionData["available"] = true

			if request.Search != "" {
				conditionData["title"] = "%" + request.Search + "%"
				conditionData["course_id"] = "%" + request.Search + "%"

				// Calculate the total number of pages based on the configured page size for given filter condition
				totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, true, []string{"=", "=", "like", "like"}, repo.db, request.PageSize)

				if totalPage < float64(request.Page) {
					request.Page = int(totalPage)
				}

				err := repo.db.Debug().Model(&models.Course{}).Where("course_group_id = ? AND available = ? AND title like ? OR course_id like ?", request.CourseGroupID, true, "%"+request.Search+"%", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

				return courses, totalPage, err
			}

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, true, []string{"=", "="}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := repo.db.Debug().Model(&models.Course{}).Where("course_group_id = ? AND available = ?", request.CourseGroupID, true).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

			return courses, totalPage, err
		}

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Debug().Model(&models.Course{}).Where("course_group_id = ?", request.CourseGroupID).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

		return courses, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"
		conditionData["course_id"] = "%" + request.Search + "%"

		if user.RoleID == 4 {
			conditionData["available"] = true

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, false, []string{"like", "like", "="}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := repo.db.Debug().Model(&models.Course{}).Where("title like ? OR course_id like ? AND available = ?", "%"+request.Search+"%", "%"+request.Search+"%", true).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

			return courses, totalPage, err
		}

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, false, []string{"like", "like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Debug().Model(&models.Course{}).Where("title like ? OR course_id like ?", "%"+request.Search+"%", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

		return courses, totalPage, err
	}

	if user.RoleID == 4 {
		conditionData := make(map[string]interface{})
		conditionData["available"] = true
		if request.Search != "" {
			conditionData["title"] = "%" + request.Search + "%"
			conditionData["course_id"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, true, []string{"=", "=", "like", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := repo.db.Debug().Model(&models.Course{}).Where("course_group_id = ? AND available = ? AND title like ? OR course_id like ?", request.CourseGroupID, true, "%"+request.Search+"%", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

			return courses, totalPage, err
		}
		totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err = repo.db.Debug().Model(&models.Course{}).Where("available = ?", true).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

		return courses, totalPage, err
	}

	totalPage := utils.GetTotalPage(models.Course{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	err = repo.db.Debug().Model(&models.Course{}).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

	return courses, totalPage, err
}

// func (repo *repository) ListCourse(request presenter.CourseListRequest) ([]models.Course, float64, error) {
// 	var courses []models.Course

// 	user, err := repo.accountRepo.GetUserByID(request.UserID)
// 	if err != nil {
// 		return courses, 0, err
// 	}

// 	// Initialize an empty map to store condition data
// 	conditionData := make(map[string]interface{})
// 	conditionData["course_group_id"] = request.CourseGroupID
// 	conditionData["title"] = "%" + request.Search + "%"
// 	conditionData["course_id"] = "%" + request.Search + "%"
// 	if request.CourseGroupID != 0 {
// 		if user.RoleID == 4 {
// 			conditionData["available"] = true
// 		}

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, user.RoleID == 4, []string{"=", "=", "like", "like"}, repo.db, request.PageSize)

// 		if totalPage < float64(request.Page) {
// 			request.Page = int(totalPage)
// 		}
// 		fmt.Printf("request.Page: %v\n", request.Page)
// 		fmt.Printf("request.PageSize: %v\n", request.PageSize)
// 		err := repo.db.Debug().Model(&models.Course{}).Where("course_group_id = ? AND available = ? AND (title like ? OR course_id like ?)", request.CourseGroupID, conditionData["available"], conditionData["title"], conditionData["course_id"]).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

// 		return courses, totalPage, err
// 	}

// 	if request.Search != "" {
// 		if user.RoleID == 4 {
// 			conditionData["available"] = true
// 		}
// 		fmt.Println("====================")

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, user.RoleID == 4, []string{"like", "like", "="}, repo.db, request.PageSize)

// 		if totalPage < float64(request.Page) {
// 			request.Page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Course{}).Where("(title like ? OR course_id like ?) AND available = ?", conditionData["title"], conditionData["course_id"], conditionData["available"]).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

// 		return courses, totalPage, err
// 	}

// 	if user.RoleID == 4 {
// 		conditionData["available"] = true
// 	}

// 	totalPage := utils.GetTotalPage(models.Course{}, repo.db, request.PageSize)

// 	if totalPage < float64(request.Page) {
// 		request.Page = int(totalPage)
// 	}

// 	err = repo.db.Debug().Model(&models.Course{}).Where("available = ?", conditionData["available"]).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&courses).Error

// 	return courses, totalPage, err
// }
