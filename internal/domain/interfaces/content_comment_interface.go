package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type ContentCommentUsecase interface {
	CreateComment(data presenter.CommentCreateUpdateRequest) map[string]string
	ListComment(request presenter.CommentListRequest) ([]presenter.CommentListResponse, int, error)
	UpdateComment(reply models.Comment) map[string]string
	DeleteComment(id uint) error
}

type ContentCommentRepository interface {
	GetCommentByID(id uint) (models.Comment, error)
	CreateComment(data presenter.CommentCreateUpdateRequest) error
	ListComment(request presenter.CommentListRequest) ([]models.Comment, float64, error)
	UpdateComment(reply models.Comment) error
	DeleteComment(id uint) error
}
