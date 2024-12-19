package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
)

// func (repo *repository) ListSubject(page int, courseID uint, search string, pageSize int) ([]models.Subject, float64, error) {
// 	var subjects []models.Subject

// 	if courseID != 0 {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["course_id"] = courseID
// 		if search != "" {
// 			// Initialize an empty map to store condition data
// 			conditionData["title"] = "%" + search + "%"
// 			conditionData["subject_id"] = "%" + search + "%"

// 			// Calculate the total number of pages based on the configured page size for given filter condition
// 			totalPage := utils.GetTotalPageByConditionModel(models.Subject{}, conditionData, false, []string{"=", "like"}, repo.db, pageSize)

// 			if totalPage < float64(page) {
// 				page = int(totalPage)
// 			}

// 			err := repo.db.Debug().Model(&models.Subject{}).Where("title like ? OR subject_id like ?", "%"+search+"%", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Find(&subjects).Order("id").Error

// 			return subjects, totalPage, err
// 		}
// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Subject{}, conditionData, true, []string{"="}, repo.db, pageSize)

// 		if totalPage < float64(page) {
// 			page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Subject{}).Where("course_id = ?", courseID).Scopes(utils.Paginate(page, pageSize)).Find(&subjects).Order("id").Error

// 		return subjects, totalPage, err
// 	}

// 	if search != "" {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["title"] = "%" + search + "%"
// 		conditionData["subject_id"] = "%" + search + "%"

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Subject{}, conditionData, false, []string{"like", "like"}, repo.db, pageSize)

// 		if totalPage < float64(page) {
// 			page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Subject{}).Where("title like ? OR subject_id like ?", "%"+search+"%", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Find(&subjects).Order("id").Error

// 		return subjects, totalPage, err
// 	}

// 	// Calculate the total number of pages based on the configured page size
// 	totalPage := utils.GetTotalPage(models.Subject{}, repo.db, pageSize)

// 	if totalPage < float64(page) {
// 		page = int(totalPage)
// 	}

// 	err := repo.db.Debug().Model(&models.Subject{}).Scopes(utils.Paginate(page, pageSize)).Find(&subjects).Order("id").Error

// 	return subjects, totalPage, err

// }

func (repo *Repository) ListSubject(page int, courseID uint, search string, pageSize int) ([]models.Subject, float64, error) {
	var subjects []models.Subject

	baseQuery := repo.db.Debug().Model(&models.Subject{})

	if courseID != 0 {
		baseQuery = baseQuery.Where("course_id = ?", courseID)
		if search != "" {
			baseQuery = baseQuery.Where("title like ? OR subject_id like ?", "%"+search+"%", "%"+search+"%")
		}
	}

	if search != "" {
		baseQuery = baseQuery.Where("title like ? OR subject_id like ?", "%"+search+"%", "%"+search+"%")
	}

	totalPage := utils.GetTotalPage(models.Subject{}, baseQuery, pageSize)

	if totalPage < float64(page) {
		page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&subjects).Error

	return subjects, totalPage, err
}
