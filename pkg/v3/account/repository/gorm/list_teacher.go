package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"math"
)

/*
ListTeacher is a repository method responsible for retrieving a paginated list of teachers from the
database. It calculates the total number of pages based on the configured page size and fetches the
specified page of teachers.

Parameters:
  - repo: A pointer to the repository struct, representing the data access layer for user-related
    operations. It provides access to the underlying database for retrieving the list of teachers.
  - page: An integer representing the page number of the paginated result set to be retrieved.

Returns:
  - []models.User: An array of User models representing the paginated list of teachers.
  - float64: The total number of pages in the paginated result set.
  - error: An error indicating any issues encountered during the retrieval of the teacher list.
    A nil error signifies a successful retrieval.
*/
func (repo *Repository) ListTeacher(request *presenter.TeacherListRequest) ([]presenter.UserResponse, float64, error) {
	var teachers []presenter.UserResponse

	if request.CourseID != 0 {
		var teacherIDs []uint
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = request.CourseID
		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["first_name"] = "%" + request.Search + "%"
			conditionData["middle_name"] = "%" + request.Search + "%"
			conditionData["last_name"] = "%" + request.Search + "%"
			conditionData["role_id"] = 3

			// Calculate the total number of pages based on the configured page size for given filter condition
			// totalPage := utils.GetTotalPageByConditionModel(models.User{}, conditionData, true, []string{"=", "like", "like", "like", "="}, repo.db, request.PageSize)

			// if totalPage < float64(request.Page) {
			// 	request.Page = int(totalPage)
			// }
			var count int64
			if err := repo.db.Debug().Model(&models.User{}).
				Joins("JOIN teachers ON users.id = teachers.id").
				Where("(users.first_name LIKE ? OR users.middle_name LIKE ? OR users.last_name LIKE ?) AND users.role_id = ? AND teachers.course_id = ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 3, request.CourseID).
				Count(&count).Error; err != nil {
				return teachers, 0, err
			}
			totalPage := math.Ceil(float64(count) / float64(request.PageSize))

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := repo.db.Debug().Model(&models.User{}).Where("(first_name like ? OR middle_name like ? OR last_name like ?) AND role_id = ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 3).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&teachers).Order("id").Error

			return teachers, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Teacher{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Select("id").Model(&models.Teacher{}).Where("course_id = ?", request.CourseID).Find(&teacherIDs).Error
		if err != nil {
			return teachers, 0, err
		}

		err = repo.db.Debug().Model(&models.User{}).Where("id IN ?", teacherIDs).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&teachers).Order("id").Error
		if err != nil {
			return teachers, 0, err
		}

		return teachers, totalPage, err
	}

	if request.SubjectID != 0 {
		var teacherIDs []uint
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["subject_id"] = request.SubjectID
		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["first_name"] = "%" + request.Search + "%"
			conditionData["middle_name"] = "%" + request.Search + "%"
			conditionData["last_name"] = "%" + request.Search + "%"
			conditionData["role_id"] = 3

			// Calculate the total number of pages based on the configured page size for given filter condition
			// totalPage := utils.GetTotalPageByConditionModel(models.User{}, conditionData, true, []string{"=", "like", "like", "like", "="}, repo.db, request.PageSize)

			// if totalPage < float64(request.Page) {
			// 	request.Page = int(totalPage)
			// }

			// err := repo.db.Debug().Model(&models.User{}).Where("(first_name like ? OR middle_name like ? OR last_name like ?) AND role_id = ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 3).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&teachers).Order("id").Error
			var count int64
			if err := repo.db.Debug().Model(&models.User{}).
				Joins("JOIN teachers ON users.id = teachers.id").
				Where("(users.first_name LIKE ? OR users.middle_name LIKE ? OR users.last_name LIKE ?) AND users.role_id = ? AND teachers.subject_id = ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 3, request.SubjectID).
				Count(&count).Error; err != nil {
				return teachers, 0, err
			}
			totalPage := math.Ceil(float64(count) / float64(request.PageSize))

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}
			err := repo.db.Debug().Model(&models.User{}).
				Joins("JOIN teachers ON users.id = teachers.id").
				Where("(users.first_name LIKE ? OR users.middle_name LIKE ? OR users.last_name LIKE ?) AND users.role_id = ? AND teachers.subject_id = ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 3, request.SubjectID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&teachers).Order("id").Error

			return teachers, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Teacher{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Select("id").Model(&models.Teacher{}).Where("subject_id = ?", request.SubjectID).Find(&teacherIDs).Error
		if err != nil {
			return teachers, 0, err
		}

		err = repo.db.Debug().Model(&models.User{}).Where("id IN ?", teacherIDs).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&teachers).Order("id").Error
		if err != nil {
			return teachers, 0, err
		}

		return teachers, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["first_name"] = "%" + request.Search + "%"
		conditionData["middle_name"] = "%" + request.Search + "%"
		conditionData["last_name"] = "%" + request.Search + "%"
		conditionData["role_id"] = 3

		// Calculate the total number of pages based on the configured page size for given filter condition
		// totalPage := utils.GetTotalPageByConditionModel(models.User{}, conditionData, true, []string{"like", "like", "like", "="}, repo.db, request.PageSize)
		// fmt.Printf("totalPage: %v\n", totalPage)
		// if totalPage < float64(request.Page) {
		// 	request.Page = int(totalPage)
		// }
		var count int64
		if err := repo.db.Debug().Model(&models.User{}).
			Where("(first_name LIKE ? OR middle_name LIKE ? OR last_name LIKE ?) AND role_id = ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 3).
			Count(&count).Error; err != nil {
			return teachers, 0, err
		}
		totalPage := math.Ceil(float64(count) / float64(request.PageSize))

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Debug().Model(&models.User{}).Where("(first_name like ? OR middle_name like ? OR last_name like ?) AND role_id = ?", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", 3).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&teachers).Order("id").Error

		return teachers, totalPage, err
	}

	conditionData := make(map[string]interface{})
	conditionData["role_id"] = 3

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPageByConditionModel(models.User{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	// Fetch a paginated list of categories from the database
	err := repo.db.Debug().Model(&models.User{}).Where("role_id = ?", 3).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&teachers).Order("id").Error

	return teachers, totalPage, err
}
