package gorm

import (
	"avyaas/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

func (repo *Repository) UpdateNotification(data models.Notification) error {
	_, err := repo.GetNotificationByID(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return repo.db.Where("id = ?", data.ID).Updates(&models.Notification{
		Title:            data.Title,
		NotificationType: data.NotificationType,
		Description:      data.Description,
		ScheduledDate:    data.ScheduledDate,
		Recipient:        data.Recipient,
		CourseID:         data.CourseID,
		Consumed:         data.Consumed,
	}).Error
}
