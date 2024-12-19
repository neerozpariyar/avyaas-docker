package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"time"
)

type LiveUsecase interface {
	CreateLive(data models.Live) map[string]string
	ListLive(request presenter.ListLiveRequest) (interface{}, int, error)
	UpdateLive(data models.Live) map[string]string
	DeleteLive(id uint) error
	MeetingSDKKey(key string, meetingID int64, role int, secret string) (string, error)
}
type LiveRepository interface {
	GetLiveByID(id uint) (models.Live, error)
	CreateLive(data models.Live) error
	GetEmailByID(id uint) (string, error)
	GetMeetingIDByLiveID(id uint) (int, error)
	GetLiveByMeetingID(meetingID int64) (models.Live, error)
	GetEmailByMeetingID(meetingID int64) (string, error)
	GetConflictingMeeting(email string, startTime, endDateTime time.Time) ([]*models.Live, error)
	ListLive(request presenter.ListLiveRequest) ([]models.Live, float64, error)
	UpdateLive(data models.Live) error
	DeleteLive(id uint) error
}
