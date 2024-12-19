package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type LiveGroupUsecase interface {
	CreateLiveGroup(data presenter.LiveGroupCreateUpdatePresenter) map[string]string
	ListLiveGroup(request presenter.ListLiveGroupRequest) ([]presenter.LiveGroupListResponse, int, error)
	UpdateLiveGroup(data models.LiveGroup) map[string]string
	DeleteLiveGroup(id uint) error
	// UploadRecording(liveGroupID uint, file *multipart.FileHeader) map[string]string
}

type LiveGroupRepository interface {
	GetLiveGroupByID(id uint) (models.LiveGroup, error)
	GetLiveGroupByTitle(title string) (models.LiveGroup, error)
	CreateLiveGroup(data presenter.LiveGroupCreateUpdatePresenter) map[string]string
	ListLiveGroup(request presenter.ListLiveGroupRequest) ([]models.LiveGroup, float64, error)
	UpdateLiveGroup(data models.LiveGroup) error
	DeleteLiveGroup(id uint) error
	// UploadRecording(liveGroupID uint, file *multipart.FileHeader) error
}
