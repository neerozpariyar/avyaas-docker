package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type NoticeUsecase interface {
	CreateNotice(data presenter.NoticeCreateUpdatePresenter) map[string]string
	ListNotice(req presenter.NoticeListReq) ([]presenter.NoticeListPresenter, int, error)
	DeleteNotice(id uint) (*models.Notice, error)
	UpdateNotice(requestBody presenter.NoticeCreateUpdatePresenter) (*models.Notice, map[string]string)
}

type NoticeRepository interface {
	CreateNotice(data presenter.NoticeCreateUpdatePresenter) error
	GetNoticeByID(id uint) (*models.Notice, error)
	ListNotice(req presenter.NoticeListReq) ([]models.Notice, float64, error)
	DeleteNotice(id uint) (*models.Notice, error)
	UpdateNotice(requestBody presenter.NoticeCreateUpdatePresenter) error
}
