package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
)

func (repo *Repository) ListFeedback(page int, courseID uint, pageSize int) ([]models.Feedback, float64, error) {
	var feedbacks []models.Feedback
	baseQuery := repo.db.Debug().Model(&models.Feedback{})
	if courseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = courseID

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Feedback{}, conditionData, true, []string{"="}, repo.db, pageSize)

		if totalPage < float64(page) {
			page = int(totalPage)
		}

		err := baseQuery.Where("course_id = ?", courseID).Order("updated_at").Scopes(utils.Paginate(page, pageSize)).Find(&feedbacks).Error

		return feedbacks, totalPage, err
	}
	// courseName:=repo.db.Debug().Model(&models.Course{}).Where("course_id = ?", courseID).First()

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.Feedback{}, repo.db, pageSize)

	if totalPage < float64(page) {
		page = int(totalPage)
	}

	err := baseQuery.Order("id").Scopes(utils.Paginate(page, pageSize)).Find(&feedbacks).Error

	return feedbacks, totalPage, err

}
