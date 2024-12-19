package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"math"
	"time"
)

func (repo *Repository) ListLive(request presenter.ListLiveRequest) ([]models.Live, float64, error) {
	var lives []models.Live
	var totalPage float64

	user, err := repo.accountRepo.GetUserByID(request.UserID)
	if err != nil {
		return lives, totalPage, err
	}

	if user.RoleID == 4 {
		conditionData := make(map[string]interface{})
		conditionData["user_id"] = request.UserID
		conditionData["course_id"] = request.CourseID

		err = repo.db.Debug().Model(&models.Live{}).Where("id IN (?)", repo.db.Select("live_id").Model(&models.StudentLive{}).Where("user_id = ? AND course_id = ?", request.UserID, request.CourseID)).Order("start_time").Find(&lives).Error
		if err != nil {
			return nil, totalPage, err
		}

		var freeLives []models.Live
		err = repo.db.Model(&models.Live{}).Where("is_free = ? AND course_id = ?", true, request.CourseID).Find(&freeLives).Error
		if err != nil {
			return nil, totalPage, err
		}

		lives = append(lives, freeLives...)

		totalPage = math.Ceil(float64(len(lives)) / float64(request.PageSize))

		return lives, totalPage, err
	}

	if request.Search != "" {
		conditionData := make(map[string]interface{})
		conditionData["topic"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Live{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := repo.db.Debug().Model(&models.Live{}).Where("topic like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&lives).Error

		return lives, totalPage, err
	}

	if request.LiveGroupID != 0 {
		// listMeetingsOpt := zoom.ListMeetingsOptions{
		// 	HostID: strconv.Itoa(int(liveGroupID)),
		// }

		// listMeetingsResponse, err := zoom.ListMeetings(listMeetingsOpt)

		// if err != nil {
		// 	log.Printf("Error: %+v\n\n", err)
		// }
		// log.Printf("ListMeetings: %+v\n\n", listMeetingsResponse)
		conditionData := make(map[string]interface{})
		conditionData["live_group_id"] = request.LiveGroupID

		// Update islive of all meetings to true whose current time matches the starttime
		repo.db.Debug().Model(&models.Live{}).Where("start_time <= ?", time.Now()).Where("end_date_time >= ?", time.Now()).Update("is_live", true)
		// Set isLive to false for meetings whose end time has passed
		// here start_time is in YYYY-MM-DD HH:mm:ss.SSS but duration is in mm only but needed return is YYYY-MM-DD HH:mm:ss.SSS with sum of minute and start_time
		sqlQuery := `
		UPDATE lives
		SET is_live = false
		WHERE DATE_FORMAT(DATE_ADD(start_time, INTERVAL duration MINUTE), '%Y-%m-%d %H:%i:%s') <= ?
		`
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		repo.db.Debug().Exec(sqlQuery, currentTime)
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Live{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		err := repo.db.Debug().Model(&models.Live{}).Where("live_group_id = ?", request.LiveGroupID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&lives).Order("id").Error

		return lives, totalPage, err
	}

	totalPage = utils.GetTotalPage(models.Live{}, repo.db, request.PageSize)
	// Update islive of all meetings to true whose current time matches the starttime
	// Fetch duration for each meeting from the database

	// err := repo.db.Debug().Model(&models.Live{}).Select("duration").Scopes(utils.Paginate(page, pageSize)).Find(&lives).Order("id").Scan(&duration).Error
	// endTime := time.Now().Add(duration)
	// repo.db.Debug().Model(&models.Live{}).Where("live_group_id = ?", liveGroupID).Where("starttime <= ?", endTime).Where("endtime >= ?", time.Now()).Update("islive", true)
	repo.db.Debug().Model(&models.Live{}).Where("start_time <= ?", time.Now()).Where("end_date_time >= ?", time.Now()).Update("is_live", true)

	// Set isLive to false for meetings whose end time has passed
	sqlQuery := `
		UPDATE lives
		SET is_live = false
		WHERE DATE_FORMAT(DATE_ADD(start_time, INTERVAL duration MINUTE), '%Y-%m-%d %H:%i:%s') <= ?
		`
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	repo.db.Debug().Exec(sqlQuery, currentTime)
	err = repo.db.Debug().Model(&models.Live{}).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&lives).Order("id").Error

	return lives, totalPage, err

}
