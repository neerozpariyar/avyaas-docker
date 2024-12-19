package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListNotification(data presenter.NotificationListRequest) ([]models.Notification, float64, error) {
	var notifications []models.Notification
	baseQuery := repo.db.Debug().Model(&models.Notification{})

	if data.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + data.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Notification{}, conditionData, false, []string{"like"}, repo.db, data.PageSize)

		if totalPage < float64(data.Page) {
			data.Page = int(totalPage)
		}

		err := repo.db.Debug().Model(&models.Notification{}).Where("title like ?", "%"+data.Search+"%").Scopes(utils.Paginate(data.Page, data.PageSize)).Order("id").Find(&notifications).Error

		return notifications, totalPage, err
	}

	if data.CourseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = data.CourseID
		if data.Search != "" {
			// Initialize an empty map to store condition data
			conditionData := make(map[string]interface{})
			conditionData["title"] = "%" + data.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Notification{}, conditionData, false, []string{"like", "="}, repo.db, data.PageSize)

			if totalPage < float64(data.Page) {
				data.Page = int(totalPage)
			}

			err := repo.db.Debug().Model(&models.Notification{}).Where("title like ? AND course_id = ?", "%"+data.Search+"%", data.CourseID).Scopes(utils.Paginate(data.Page, data.PageSize)).Order("id").Find(&notifications).Error

			return notifications, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Notification{}, conditionData, true, []string{"="}, repo.db, data.PageSize)

		if totalPage < float64(data.Page) {
			data.Page = int(totalPage)
		}

		err := baseQuery.Where("course_id = ?", data.CourseID).Scopes(utils.Paginate(data.Page, data.PageSize)).Order("updated_at").Find(&notifications).Error

		return notifications, totalPage, err
	}

	// if data.UserID != 0 {
	// 	// Initialize an empty map to store condition data
	// 	conditionData := make(map[string]interface{})
	// 	conditionData["user_id"] = data.UserID

	// 	// Calculate the total number of pages based on the configured page size for given filter condition
	// 	totalPage := utils.GetTotalPageByConditionModel(models.Notification{}, conditionData, true, []string{"="}, repo.db, data.PageSize)

	// 	if totalPage < float64(data.Page) {
	// 		data.Page = int(totalPage)
	// 	}

	// 	err := baseQuery.Where("user_id = ?", data.UserID).Scopes(utils.Paginate(data.Page, data.PageSize)).Order("updated_at").Find(&notifications).Error

	// 	return notifications, totalPage, err
	// }
	// courseName:=repo.db.Debug().Model(&models.Course{}).Where("course_id = ?", courseID).First()

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.Notification{}, repo.db, data.PageSize)

	if totalPage < float64(data.Page) {
		data.Page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(data.Page, data.PageSize)).Order("id").Find(&notifications).Error

	return notifications, totalPage, err

}
