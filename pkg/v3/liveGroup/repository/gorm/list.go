package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListLiveGroup(request presenter.ListLiveGroupRequest) ([]models.LiveGroup, float64, error) {
	var liveGroups []models.LiveGroup
	baseQuery := repo.db.Debug().Model(&models.LiveGroup{}).Order("id")

	if request.Search != "" {
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.LiveGroup{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&liveGroups).Error

		return liveGroups, totalPage, err
	}

	if request.CourseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = request.CourseID

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.LiveGroup{}, conditionData, true, []string{"="}, baseQuery, request.PageSize)

		err := baseQuery.Where("course_id = ?", request.CourseID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&liveGroups).Error

		return liveGroups, totalPage, err
	}

	totalPage := utils.GetTotalPage(models.LiveGroup{}, baseQuery, request.PageSize)

	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&liveGroups).Error

	return liveGroups, totalPage, err
}
