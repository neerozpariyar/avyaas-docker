package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type BookmarkUsecase interface {
	CreateBookmark(data presenter.BookmarkCreateUpdateRequest) map[string]string
	ListBookmark(data presenter.BookmarkListRequest) ([]presenter.BookmarkListResponse, int, error)
	GetBookmarkDetails(id uint) (*presenter.BookmarkDetailResponse, map[string]string)
	// UpdateBookmark(data presenter.BookmarkCreateUpdateRequest) map[string]string
	DeleteBookmark(id uint) error
}

type BookmarkRepository interface {
	GetBookmarkByID(id uint) (models.Bookmark, error)
	GetBookmarkTitleByID(id uint) (string, error)
	CreateBookmark(data presenter.BookmarkCreateUpdateRequest) error
	ListBookmark(data presenter.BookmarkListRequest) ([]models.Bookmark, float64, error)
	// GetBookmarkDetails(id uint) (*presenter.BookmarkDetailResponse, error)
	GetBookmarkTypeByID(id uint) (string, error)
	GetBookmarkedContentAndCheckIfBookmarked(userID, contentID uint) (models.Bookmark, bool, error)
	GetBookmarkedQuestionAndCheckIfBookmarked(userID, questionID uint) (models.Bookmark, bool, error)

	// UpdateBookmark(data presenter.BookmarkCreateUpdateRequest) error
	DeleteBookmark(id uint) error
}
