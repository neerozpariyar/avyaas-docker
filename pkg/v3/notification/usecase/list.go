package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListNotification(request presenter.NotificationListRequest) ([]presenter.NotificationListResponse, int, error) {
	notifications, totalPage, err := u.repo.ListNotification(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allNotifications []presenter.NotificationListResponse
	for _, notification := range notifications {
		eachNotification := presenter.NotificationListResponse{
			ID:               notification.ID,
			Title:            notification.Title,
			Description:      notification.Description,
			NotificationType: notification.NotificationType,
			Recipient:        notification.Recipient,
		}

		if notification.CourseID != 0 {
			course, err := u.courseRepo.GetCourseByID(notification.CourseID)
			if err != nil {
				return nil, 0, err
			}

			courseData := make(map[string]interface{})
			courseData["id"] = course.ID
			courseData["courseID"] = course.CourseID
			eachNotification.Course = courseData
		}

		// Format the start time and end time to UTC string format
		if notification.ScheduledDate != nil {
			eachNotification.ScheduledDate = notification.ScheduledDate.UTC().Format("2006-01-02T15:04:05Z")
		}

		allNotifications = append(allNotifications, eachNotification)
	}

	return allNotifications, int(totalPage), nil
}
