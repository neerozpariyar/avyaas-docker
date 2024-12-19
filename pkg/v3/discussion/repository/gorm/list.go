package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListDiscussion(request presenter.DiscussionListRequest) ([]models.Discussion, float64, error) {
	var discussions []models.Discussion

	// user, err := repo.accountRepo.GetUserByID(request.UserID)
	// if err != nil {
	// 	return discussions, 0, err
	// }
	//var replies []models.Reply
	if request.SubjectID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["subject_id"] = request.SubjectID
		if request.Search != "" {
			// Initialize an empty map to store condition data
			conditionData["title"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Discussion{}, conditionData, false, []string{"=", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := repo.db.Debug().Exec(`update discussions set reply_count =
			(select COUNT(id) from replies where replies.discussion_id = discussions.Id)`).Model(&models.Discussion{}).Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&discussions).Error

			return discussions, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Discussion{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Debug().Exec(`update discussions set reply_count =
		(select COUNT(id) from replies where replies.discussion_id = discussions.Id)`).Model(&models.Discussion{}).Where("subject_id = ?", request.SubjectID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&discussions).Error

		return discussions, totalPage, err
	}

	if request.Search != "" {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Discussion{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Debug().Exec(`update discussions set reply_count =
		(select COUNT(id) from replies where replies.discussion_id = discussions.Id)`).Model(&models.Discussion{}).Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&discussions).Error

		return discussions, totalPage, err
	}
	// if user.RoleID == 4 {
	// 	conditionData := make(map[string]interface{})
	// 	conditionData["available"] = true

	// 	totalPage := utils.GetTotalPageByConditionModel(models.Course{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

	// 	if totalPage < float64(request.Page) {
	// 		request.Page = int(totalPage)
	// 	}

	// 	err = repo.db.Debug().Model(&models.Discussion{}).
	// 		Scopes(utils.Paginate(request.Page, request.PageSize)).
	// 		Order("id").
	// 		Select("discussions.*, votes.has_liked").
	// 		Joins("LEFT JOIN votes ON discussions.id = votes.discussion_id AND votes.user_id = ?", user.ID).
	// 		Find(&discussions).Error

	// 	return discussions, totalPage, err
	// }
	totalPage := utils.GetTotalPage(models.Discussion{}, repo.db, request.PageSize)

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}
	//	err := repo.db.Debug().Where("id=8").Model(&models.Discussion{}).Scopes(utils.Paginate(request.Page)).Order("id").Preload("Reply").Find(&discussions).Error

	// err := repo.db.Debug().Model(&models.Discussion{}).Where("id = ?", request.DiscussionID).Scopes(utils.Paginate(request.Page)).Order("id").Preload("Reply").Find(&discussions).Error
	// fmt.Printf("len(discussions): %v\n", len(replies))

	//err := repo.db.Exec(`UPDATE discussions SET vote_count = (SELECT COUNT(vote) FROM discussions WHERE discussions.Id = ?)`, &models.Discussion{}).Error

	err := repo.db.Debug().Exec(`update discussions set reply_count =
		(select COUNT(id) from replies where replies.discussion_id = discussions.Id)`).Model(&models.Discussion{}).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&discussions).Error

	// err := repo.db.Debug().Exec(`update discussions as d,(SELECT d.Id as Id, COUNT(replies.discussion_id) AS c FROM discussions AS d
	// LEFT JOIN replies AS r ON d.Id=r.discussion_id GROUP BY d.Id) AS dc SET d.ReplyCount=dc.c WHERE d.Id=dc.Id`).Model(&models.Discussion{}).Scopes(utils.Paginate(request.Page)).Order("id").Find(&discussions).Error

	// err := repo.db.Debug().Set("reply_count", &models.Discussion{}).Count(&models.rep).Select(&models.Reply{}).Where("replies.discussion_id = discussions.Id").Error
	return discussions, totalPage, err

}
