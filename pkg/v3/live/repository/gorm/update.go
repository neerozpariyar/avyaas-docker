package gorm

import (
	"avyaas/internal/domain/models"
	"log"
)

func (repo *Repository) UpdateLive(data models.Live) error {
	transaction := repo.db.Begin()

	// err := repo.db.Select("meeting_id").First(&models.Live{}).Where("id=?", data.ID).Debug().Error
	// if err != nil {
	// 	return err
	// }
	// if metID!=data.MeetingID{
	// 	return fmt.Errorf("")
	// }
	err := transaction.Debug().
		Where("id = ?", data.ID).
		Updates(&models.Live{
			Topic:       data.Topic,
			StartTime:   data.StartTime,
			EndDateTime: data.EndDateTime,
			Duration:    data.Duration,
		}).Error

	if err != nil {
		transaction.Rollback()
		log.Printf("Error updating local database record: %+v\n", err)
		return err
	}
	transaction.Commit()

	return nil
}
