package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type ReplyUsecase interface {
	CreateReply(data presenter.ReplyCreateUpdateRequest) map[string]string
	ListReply(request presenter.ReplyListRequest) ([]presenter.ReplyListResponse, int, error)
	UpdateReply(reply models.Reply) map[string]string
	DeleteReply(id uint) error
}

type ReplyRepository interface {
	GetReplyByID(id uint) (models.Reply, error)
	CreateReply(data presenter.ReplyCreateUpdateRequest) error
	ListReply(request presenter.ReplyListRequest) ([]models.Reply, float64, error)
	UpdateReply(reply models.Reply) error
	DeleteReply(id uint) error
}
