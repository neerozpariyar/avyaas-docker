package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type ChapterUsecase interface {
	CreateChapter(chapter models.Chapter) map[string]string
	ListChapter(data presenter.ChapterListRequest) ([]presenter.Chapter, int, error)
	UpdateChapter(chapter models.Chapter) map[string]string
	DeleteChapter(id uint) error

	UpdateChapterPosition(data presenter.UpdateChapterPositionRequest) map[string]string
}

type ChapterRepository interface {
	//	GetChapterByChapterID(ChapterID string) (models.Chapter, error)
	GetChapterByID(id uint) (models.Chapter, error)

	CreateChapter(chapter models.Chapter) error
	ListChapter(data presenter.ChapterListRequest) ([]models.Chapter, float64, error)
	UpdateChapter(chapter models.Chapter) error
	DeleteChapter(id uint) error

	UpdateChapterPosition(data presenter.UpdateChapterPositionRequest) map[string]string
}
