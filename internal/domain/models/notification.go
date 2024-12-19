package models

import "time"

type Notification struct {
	Timestamp
	Title            string     `json:"title"`
	Description      string     `json:"description"`      //can include external links too
	NotificationType string     `json:"notificationType"` //announcement and push notification
	ScheduledDate    *time.Time `json:"scheduledDate"`    //publishes for the scheduled date  and time
	Recipient        string     `json:"recipient"`        //verified user, unverified user and in the specific courses
	CourseID         uint       `json:"courseID"`
	Consumed         bool       `json:"consumed"` //turns true when the notification is consumed, not published
}
