package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (repo *Repository) GetConflictingMeeting(email string, startTime, endDateTime time.Time) ([]*models.Live, error) {

	var conflictingMeetings []*models.Live
	// SELECT * FROM `lives` WHERE (email = 'abc@gmail.com' AND start_time >= '2024-02-17 20:00:05' AND end_date_time <= '2024-03-17 18:32:05')ORDER BY `lives`.`id`;
	err := repo.db.Where("email = ? AND start_time >= ? AND end_date_time <= ?", email, startTime, endDateTime).Find(&conflictingMeetings).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return conflictingMeetings, nil
}
