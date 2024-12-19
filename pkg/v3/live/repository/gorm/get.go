package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetLiveByID(id uint) (models.Live, error) {
	var live models.Live

	err := repo.db.Where("id = ?", id).First(&live).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Live{}, fmt.Errorf("live with live id: '%d' not found", id)
		}

		return models.Live{}, err
	}
	return live, nil
}

func (repo *Repository) GetMeetingIDByLiveID(id uint) (int, error) {
	var live models.Live
	err := repo.db.Where("id=?", id).First(&live).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return live.MeetingID, fmt.Errorf("live with live id: '%d' not found", id)
		}

		return live.MeetingID, err
	}
	return live.MeetingID, nil
}
func (repo *Repository) GetLiveByMeetingID(meetingID int64) (models.Live, error) {
	var live models.Live

	err := repo.db.Where("meeting_id = ?", meetingID).First(&live).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Live{}, fmt.Errorf("live with live id: '%d' not found", meetingID)
		}

		return models.Live{}, err
	}
	return live, nil
}
