package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"math"
)

// func (repo *repository) ListUnit(page int, subjectID uint, search string, pageSize int) ([]models.Unit, float64, error) {
// 	var units []models.Unit
// 	if subjectID != 0 {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["subject_id"] = subjectID
// 		if search != "" {
// 			// Initialize an empty map to store condition data
// 			conditionData["title"] = "%" + search + "%"

// 			// Calculate the total number of pages based on the configured page size for given filter condition
// 			totalPage := utils.GetTotalPageByConditionModel(models.Unit{}, conditionData, false, []string{"=", "like"}, repo.db, pageSize)

// 			if totalPage < float64(page) {
// 				page = int(totalPage)
// 			}

// 			err := repo.db.Debug().Model(&models.Unit{}).Where("title like ?", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&units).Error

// 			return units, totalPage, err
// 		}
// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Unit{}, conditionData, true, []string{"="}, repo.db, pageSize)

// 		if totalPage < float64(page) {
// 			page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Unit{}).Where("subject_id = ?", subjectID).Scopes(utils.Paginate(page, pageSize)).Order("position").Find(&units).Error

// 		return units, totalPage, err
// 	}

// 	if search != "" {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["title"] = "%" + search + "%"

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Unit{}, conditionData, false, []string{"like"}, repo.db, pageSize)

// 		if totalPage < float64(page) {
// 			page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Unit{}).Where("title like ?", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&units).Error

// 		return units, totalPage, err
// 	}

// 	// Calculate the total number of pages based on the configured page size
// 	totalPage := utils.GetTotalPage(models.Unit{}, repo.db, pageSize)

// 	if totalPage < float64(page) {
// 		page = int(totalPage)
// 	}

// 	err := repo.db.Debug().Model(&models.Unit{}).Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&units).Error

// 	return units, totalPage, err

// }

func (repo *Repository) ListUnit(page int, subjectID uint, search string, pageSize int) ([]models.Unit, float64, error) {
	var units []models.Unit

	baseQuery := repo.db.Debug().Model(&models.Unit{})

	if subjectID != 0 {
		baseQuery = baseQuery.Joins("JOIN unit_chapter_contents ucc ON units.id = ucc.unit_id").
			Joins("LEFT JOIN subject_unit_chapter_contents succ ON ucc.id = succ.unit_chapter_content_id").
			Where(" succ.subject_id = ?", subjectID)

	}

	if search != "" {
		baseQuery = baseQuery.Where("title like ? ", "%"+search+"%")
	}

	err := baseQuery.Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&units).Error

	if err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(len(units)) / float64(pageSize))

	return units, totalPage, err

}
