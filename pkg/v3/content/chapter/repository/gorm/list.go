package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"math"
)

// func (repo *repository) ListChapter(page int, unitID uint, search string, pageSize int) ([]models.Chapter, float64, error) {
// 	var chapters []models.Chapter

// 	if unitID != 0 {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["unit_id"] = unitID
// 		if search != "" {
// 			// Initialize an empty map to store condition data
// 			conditionData["title"] = "%" + search + "%"

// 			// Calculate the total number of pages based on the configured page size for given filter condition
// 			totalPage := utils.GetTotalPageByConditionModel(models.Chapter{}, conditionData, false, []string{"=", "like"}, repo.db, pageSize)

// 			if totalPage < float64(page) {
// 				page = int(totalPage)
// 			}

// 			err := repo.db.Debug().Model(&models.Chapter{}).Where("title like ?", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&chapters).Error

// 			return chapters, totalPage, err
// 		}
// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Chapter{}, conditionData, true, []string{"="}, repo.db, pageSize)

// 		if totalPage < float64(page) {
// 			page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Chapter{}).Where("unit_id = ?", unitID).Scopes(utils.Paginate(page, pageSize)).Order("position").Find(&chapters).Error

// 		return chapters, totalPage, err
// 	}

// 	if search != "" {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["title"] = "%" + search + "%"

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Chapter{}, conditionData, false, []string{"like"}, repo.db, pageSize)

// 		if totalPage < float64(page) {
// 			page = int(totalPage)
// 		}

// 		err := repo.db.Debug().Model(&models.Chapter{}).Where("title like ?", "%"+search+"%").Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&chapters).Error

// 		return chapters, totalPage, err
// 	}

// 	totalPage := utils.GetTotalPage(models.Chapter{}, repo.db, pageSize)

// 	if totalPage < float64(page) {
// 		page = int(totalPage)
// 	}

// 	err := repo.db.Debug().Model(&models.Chapter{}).Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&chapters).Error

// 	return chapters, totalPage, err
// }

func (repo *Repository) ListChapter(data presenter.ChapterListRequest) ([]models.Chapter, float64, error) {
	var chapters []models.Chapter

	baseQuery := repo.db.Debug().Model(&models.Chapter{}).Scopes(utils.Paginate(data.Page, data.PageSize))

	if data.ChapterFilter.SubjectID != 0 {
		baseQuery = baseQuery.Joins("JOIN unit_chapter_contents ucc ON chapters.id = ucc.chapter_id").
			Joins("LEFT JOIN subject_unit_chapter_contents succ ON ucc.id = succ.unit_chapter_content_id").
			Where("ucc.unit_id = ?  AND succ.subject_id = ?", data.ChapterFilter.UnitID, data.ChapterFilter.SubjectID)
	}

	if data.Search != "" {
		baseQuery = baseQuery.Where("title like ?", "%"+data.Search+"%").Order("id")
	}

	err := baseQuery.Find(&chapters).Error

	if err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(len(chapters)) / float64(data.PageSize))

	return chapters, totalPage, err
}
