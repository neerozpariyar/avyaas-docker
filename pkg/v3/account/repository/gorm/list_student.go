package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"math"
)

func (repo *Repository) ListStudent(request *presenter.StudentListRequest) ([]presenter.UserResponse, float64, error) {
	var students []presenter.UserResponse
	var err error

	if request.CourseID != 0 {
		var studentIDs []uint
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = request.CourseID
		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["first_name"] = "%" + request.Search + "%"
			conditionData["middle_name"] = "%" + request.Search + "%"
			conditionData["last_name"] = "%" + request.Search + "%"
			conditionData["email"] = "%" + request.Search + "%"
			conditionData["phone"] = "%" + request.Search + "%"
			conditionData["username"] = "%" + request.Search + "%"
			conditionData["role_id"] = 4

			// // Calculate the total number of pages based on the configured page size for given filter condition
			// totalPage := utils.GetTotalPageByConditionModel(models.User{}, conditionData, true, []string{"=", "like", "like", "like", "="}, repo.db, request.PageSize)
			// if totalPage < float64(request.Page) {
			// 	request.Page = int(totalPage)
			// }
			var count int64
			if err := repo.db.Debug().Model(&models.User{}).
				Joins("JOIN student_courses ON users.id = student_courses.user_id").
				Where("(users.first_name LIKE ? OR users.middle_name LIKE ? OR users.last_name LIKE ? AND users.email LIKE ? OR users.phone LIKE ? OR users.username LIKE ?) AND users.role_id = ? AND student_courses.course_id = ?",
					"%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 4, request.CourseID).
				Find(&students).Count(&count).Error; err != nil {
				return students, 0, err
			}

			totalPage := math.Ceil(float64(count) / float64(request.PageSize))

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			// err := repo.db.Debug().Model(&models.User{}).Where("first_name like ? OR middle_name like ? OR last_name like ? AND role_id like ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 4).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&students).Order("id").Error

			return students, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.StudentCourse{}, conditionData, true, []string{"="}, repo.db, request.Page)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Select("user_id").Model(&models.StudentCourse{}).Where("course_id = ?", request.CourseID).Find(&studentIDs).Error
		if err != nil {
			return students, 0, err
		}

		err = repo.db.Debug().Model(&models.User{}).Where("id IN ?", studentIDs).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&students).Order("id").Error
		if err != nil {
			return students, 0, err
		}

		return students, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["first_name"] = "%" + request.Search + "%"
		conditionData["middle_name"] = "%" + request.Search + "%"
		conditionData["last_name"] = "%" + request.Search + "%"
		conditionData["email"] = "%" + request.Search + "%"
		conditionData["phone"] = "%" + request.Search + "%"
		conditionData["username"] = "%" + request.Search + "%"
		conditionData["role_id"] = 4

		// Calculate the total number of pages based on the configured page size for given filter condition
		// totalPage := utils.GetTotalPageByConditionModel(models.User{}, conditionData, true, []string{"like", "like", "like", "like", "like", "="}, repo.db, request.PageSize)

		// if totalPage < float64(request.Page) {
		// 	request.Page = int(totalPage)
		// }

		err := repo.db.Debug().Model(&models.User{}).Where("(first_name like ? OR middle_name like ? OR last_name like ? OR email like ? OR phone like ? OR username like ?) AND role_id = ?",
			"%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 4).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&students).Order("id").Error

		totalPage := math.Ceil(float64(len(students)) / float64(request.PageSize))

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		return students, totalPage, err
	}

	conditionData := make(map[string]interface{})
	conditionData["role_id"] = 4

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPageByConditionModel(models.User{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	// Fetch a paginated list of categories from the database
	err = repo.db.Debug().Model(&models.User{}).Where("role_id = ?", 4).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&students).Order("id").Error

	return students, totalPage, err
}
