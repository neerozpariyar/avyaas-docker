package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListReferral(request presenter.ReferralListRequest) ([]models.Referral, float64, error) {
	var referrals []models.Referral
	baseQuery := repo.db.Debug().Model(&models.Referral{}).Order("id")

	if request.CourseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		// conditionData["subject_id"] = request.CourseID

		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["title"] = "%" + request.Search + "%"
			conditionData["code"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Referral{}, conditionData, false, []string{"=", "like", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("title like ? OR code like ?", "%"+request.Search+"%", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&referrals).Error

			return referrals, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Referral{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("course_id = ?", request.CourseID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&referrals).Error

		return referrals, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"
		conditionData["code"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Referral{}, conditionData, false, []string{"like", "like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("title like ? OR code like ?", "%"+request.Search+"%", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&referrals).Error

		return referrals, totalPage, err
	}

	totalPage := utils.GetTotalPage(models.Referral{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	err := baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&referrals).Error

	return referrals, totalPage, err
}
