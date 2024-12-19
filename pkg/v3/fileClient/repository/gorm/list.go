package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListObjects(req *presenter.FileListReq) ([]presenter.FileListRes, float64, error) {
	var files []presenter.FileListRes
	baseQuery := repo.db.Debug().Model(&models.File{})

	if !req.IsActive { //enters this block only if isActive is false
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["is_active"] = req.IsActive
		if req.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["title"] = "%" + req.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.File{}, conditionData, false, []string{"=", "like"}, repo.db, req.PageSize)

			if totalPage < float64(req.Page) {
				req.Page = int(totalPage)
			}

			err := baseQuery.Where("is_active = ? AND title like ? ", req.IsActive, "%"+req.Search+"%").Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&files).Order("id").Error

			return files, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.File{}, conditionData, true, []string{"="}, repo.db, req.PageSize)

		if totalPage < float64(req.Page) {
			req.Page = int(totalPage)
		}

		err := baseQuery.Where("is_active = ?", req.IsActive).Scopes(utils.Paginate(req.Page, req.PageSize)).Order("updated_at").Find(&files).Error

		return files, totalPage, err
	}
	if req.Service != "" { //here service means name of the folder
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["url"] = "%" + req.Service + "%"

		if req.Search != "" {

			// Initialize an empty map to store condition data
			conditionData["title"] = "%" + req.Search + "%"
			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.File{}, conditionData, false, []string{"like", "like"}, repo.db, req.PageSize)

			if totalPage < float64(req.Page) {
				req.Page = int(totalPage)
			}

			err := baseQuery.Where("title like ? ", "%"+req.Search+"%").Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&files).Order("id").Error

			return files, totalPage, err
		}

		err := baseQuery.Scopes(utils.Paginate(req.Page, req.PageSize)).Where("url LIKE ?", "%"+req.Service+"%").Order("id").Find(&files).Error

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.File{}, conditionData, true, []string{"like"}, repo.db, req.PageSize)

		if totalPage < float64(req.Page) {
			req.Page = int(totalPage)
		}
		return files, totalPage, err
	}
	if req.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + req.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.File{}, conditionData, false, []string{"like", "like"}, repo.db, req.PageSize)

		if totalPage < float64(req.Page) {
			req.Page = int(totalPage)
		}

		err := baseQuery.Where("title like ? ", "%"+req.Search+"%").Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&files).Order("id").Error
		return files, totalPage, err
	}
	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.File{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(req.Page, req.PageSize)).Order("id").Find(&files).Error
	return files, totalPage, err

}
