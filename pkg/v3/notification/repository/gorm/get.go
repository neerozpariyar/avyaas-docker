package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetNotificationByID(id uint) (models.Notification, error) {
	var notification models.Notification

	// Retrieve the notification from the database based on given id
	err := repo.db.Where("id = ?", id).First(&notification).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Notification{}, fmt.Errorf("notification with notification id: '%d' not found", id)
		}

		return models.Notification{}, err
	}

	return notification, nil
}
