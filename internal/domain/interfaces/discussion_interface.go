package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type DiscussionUsecase interface {
	CreateDiscussion(data presenter.DiscussionCreateUpdateRequest) map[string]string
	ListDiscussion(request presenter.DiscussionListRequest) ([]presenter.Discussion, int, error)
	UpdateDiscussion(discussion presenter.DiscussionCreateUpdateRequest) map[string]string
	DeleteDiscussion(id uint) error
	LikeOrUnlikeDiscussion(userID, discussionID uint) map[string]string
}

type DiscussionRepository interface {
	GetDiscussionByID(id uint) (models.Discussion, error)
	CreateDiscussion(data presenter.DiscussionCreateUpdateRequest) error
	ListDiscussion(request presenter.DiscussionListRequest) ([]models.Discussion, float64, error)
	UpdateDiscussion(discussion presenter.DiscussionCreateUpdateRequest) error
	DeleteDiscussion(id uint) error
	LikeOrUnlikeDiscussion(userID, discussionID uint) error
	GetHasLikedValue(discussionID, userID uint) (bool, error)
}
