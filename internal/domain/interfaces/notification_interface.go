package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type NotificationUsecase interface {
	CreateNotification(data models.Notification) map[string]string
	ListNotification(request presenter.NotificationListRequest) ([]presenter.NotificationListResponse, int, error)
	UpdateNotification(data models.Notification) map[string]string
	DeleteNotification(id uint) error
	PublishNotification(notificationID uint) error
	// ConsumeNotification()
	AddFCMToken(fcmToken models.FCMToken) (err error)
	GetFCMByUserIDAndToken(userID uint, token string) (fcmToken models.FCMToken, err error)
	FCMRegister(fcmToken models.FCMToken) error
}

type NotificationRepository interface {
	GetNotificationByID(id uint) (models.Notification, error)
	// PublishNotification(notificationID uint) error
	CreateNotification(data models.Notification) error
	ListNotification(request presenter.NotificationListRequest) ([]models.Notification, float64, error)
	GetUsersEnrolledInCourse(courseID uint) ([]models.User, error)
	UpdateNotification(data models.Notification) error
	DeleteNotification(id uint) error

	AddFCMToken(fcmToken models.FCMToken) (err error)
	GetFCMByUserIDAndToken(userID uint, token string) (fcmToken models.FCMToken, err error)

	GetVerifiedUsers() ([]models.User, error)
	GetUnverifiedUsers() ([]models.User, error)
	GetUsersByCourseID(courseID uint) ([]models.User, error)
}
