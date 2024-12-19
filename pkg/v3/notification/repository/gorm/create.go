package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) CreateNotification(data models.Notification) error {
	transaction := repo.db.Begin()
	// var err error
	// var recipient string

	// Check if the recipient is a course
	// if strings.EqualFold(data.Recipient, "course") {
	// 	// Fetch the title of the course
	// 	var courseTitle string
	// 	err = repo.db.Model(&models.Course{}).Select("title").Where("id = ?", data.CourseID).Scan(&courseTitle).Error
	// 	if err != nil {
	// 		transaction.Rollback()
	// 		return err
	// 	}
	// 	recipient = courseTitle
	// } else if strings.EqualFold(data.Recipient, "verified") || strings.EqualFold(data.Recipient, "unverified") {
	// 	// If the recipient is "verified" or "unverified", set it directly
	// 	recipient = data.Recipient
	// } else {
	// 	// If the recipient is none of the above, return an error
	// 	transaction.Rollback()
	// 	return errors.New("invalid recipient type")
	// }

	err := repo.db.Create(&models.Notification{
		Title:            data.Title,
		ScheduledDate:    data.ScheduledDate,
		Description:      data.Description,
		NotificationType: data.NotificationType,
		Recipient:        data.Recipient,
		CourseID:         data.CourseID,
		Consumed:         false,
	}).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
