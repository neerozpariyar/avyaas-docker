package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

// func (repo *repository) ListNote(page int, contentID uint, pageSize int) ([]models.Note, float64, error) {
// 	var notes []models.Note
// 	if contentID != 0 {
// 		// Initialize an empty map to store condition data
// 		conditionData := make(map[string]interface{})
// 		conditionData["content_id"] = contentID

// 		// Calculate the total number of pages based on the configured page size for given filter condition
// 		totalPage := utils.GetTotalPageByConditionModel(models.Note{}, conditionData, true, []string{"="}, repo.db, pageSize)

// 		err := repo.db.Debug().Model(&models.Note{}).Where("content_id = ?", contentID).Scopes(utils.Paginate(page, pageSize)).Find(&notes).Order("id").Error

// 		return notes, totalPage, err
// 	}

// 	totalPage := utils.GetTotalPage(models.Note{}, repo.db, pageSize)

// 	err := repo.db.Debug().Model(&models.Note{}).Scopes(utils.Paginate(page, pageSize)).Find(&notes).Order("id").Error

// 	return notes, totalPage, err

// }

func (repo *Repository) ListNote(data presenter.NoteListRequest) ([]models.Note, float64, error) {
	var notes []models.Note
	var totalPage float64
	var err error

	baseQuery := repo.db.Debug().Model(&models.Note{}).Scopes(utils.Paginate(data.Page, data.PageSize)).Order("id")

	if data.ContentID != 0 {
		baseQuery = baseQuery.Where("content_id = ?", data.ContentID)
		conditionData := map[string]interface{}{"content_id": data.ContentID}
		totalPage = utils.GetTotalPageByConditionModel(models.Note{}, conditionData, true, []string{"="}, repo.db, data.PageSize)
		if data.Search != "" {
			baseQuery = baseQuery.Where("title LIKE ? AND content_id = ?", "%"+data.Search+"%", data.ContentID)
		}
	} else {
		totalPage = utils.GetTotalPage(models.Note{}, repo.db, data.PageSize)
		if data.Search != "" {
			baseQuery = baseQuery.Where("title LIKE ?", "%"+data.Search+"%")
		}
	}

	err = baseQuery.Find(&notes).Error

	return notes, totalPage, err
}
