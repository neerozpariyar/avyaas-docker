package presenter

import (
	"avyaas/internal/domain/models"
	"time"
)

type NotificationListRequest struct {
	Page     int
	PageSize int
	CourseID uint
	UserID   uint
	Search   string
}

type NotificationCreateUpdateRequest struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`      //can include external links too
	NotificationType string    `json:"notificationType"` //announcement and push notification
	ScheduledDate    time.Time `json:"scheduledDate"`    //publishes for the scheduled date  and time
	CourseID         uint      `json:"courseID"`
}

type NotificationListResponse struct {
	ID               uint        `json:"id"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	NotificationType string      `json:"notificationType"`
	ScheduledDate    string      `json:"scheduledDate"`
	Course           interface{} `json:"course"`
	Recipient        string      `json:"recipient"`
}

type ConsumerResponse struct {
	Recipient string                `json:"recipient"`
	Data      []models.Notification `json:"data"`
}
