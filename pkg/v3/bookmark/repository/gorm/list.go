package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListBookmark(data presenter.BookmarkListRequest) ([]models.Bookmark, float64, error) {
	var bookmarks []models.Bookmark

	// Initialize an empty map to store condition data
	conditionData := make(map[string]interface{})
	conditionOperators := []string{"="}

	if data.BookmarkType == "content" || data.BookmarkType == "question" {
		conditionData["bookmark_type"] = data.BookmarkType
	}
	
	// Add condition to filter bookmarks by the logged in user
	if data.UserID != 0 {
		conditionData["user_id"] = data.UserID
		conditionOperators = append(conditionOperators, "=")
	}

	if data.Search != "" {
		conditionData["title"] = "%" + data.Search + "%"
		conditionOperators = append(conditionOperators, "like")
		totalPage := utils.GetTotalPageByConditionModel(models.Bookmark{}, conditionData, len(conditionData) > 0, conditionOperators, repo.db, data.PageSize)

		if totalPage < float64(data.Page) {
			data.Page = int(totalPage)
		}
	}

	query := repo.db.Debug().Model(&models.Bookmark{}).Scopes(utils.Paginate(data.Page, data.PageSize))
	// Add the conditions to the query
	for field, value := range conditionData {
		if field == "title" {
			query = query.Where(field+" like ?", value)
		} else {
			query = query.Where(field+" = ?", value)
		}
	}

	err := query.Scopes(utils.Paginate(data.Page, data.PageSize)).Order("updated_at").Find(&bookmarks).Error
	totalPage := utils.GetTotalPage(models.Bookmark{}, repo.db, data.PageSize)
	if totalPage < float64(data.Page) {
		data.Page = int(totalPage)
	}

	return bookmarks, totalPage, err
}
