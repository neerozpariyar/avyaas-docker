package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type RecordingUsecase interface {
	UploadRecording(data presenter.RecordingCreateUpdateRequest) map[string]string
	ListRecording(request presenter.RecordingListRequest) ([]models.Recording, int, error)
	UpdateRecording(data presenter.RecordingCreateUpdateRequest) map[string]string
	DeleteRecording(id uint) error
}
type RecordingRepository interface {
	UploadRecording(data presenter.RecordingCreateUpdateRequest) error
	ListRecording(request presenter.RecordingListRequest) ([]models.Recording, float64, error)
	GetRecordingByID(id uint) (models.Recording, error)
	UpdateRecording(data presenter.RecordingCreateUpdateRequest) error
	DeleteRecording(id uint) error
}
