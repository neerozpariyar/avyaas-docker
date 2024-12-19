package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type ContentUsecase interface {
	CreateContent(data presenter.ContentCreateUpdateRequest) map[string]string
	ListContent(request presenter.ContentListRequest) ([]presenter.SingleContentResponse, int, error)
	UpdateContent(data presenter.ContentCreateUpdateRequest) map[string]string
	DeleteContent(id uint) error
	GetContentDetails(id, requesterID uint) (*presenter.ContentDetailResponse, map[string]string)
	AssignContentsToChapter(data models.ChapterContent) error
	UpdateContentPosition(data presenter.UpdateContentPositionRequest) map[string]string

	// progress
	MarkAsCompleted(userID, contentID uint) map[string]string
	EvaluateProgress(data presenter.ProgressPresenter) error
	GetCourseIDByContentID(contentID uint) (uint, error)
	EvaluateContentProgress(data presenter.ContentProgressPresenter) error
	AssignContentsToRelation(uint, uint, []uint) map[string]string
}

type ContentRepository interface {
	CreateContent(data presenter.ContentCreateUpdateRequest) error
	ListContent(request presenter.ContentListRequest) ([]models.Content, float64, error)
	GetContentByID(id uint) (models.Content, error)
	UpdateContent(data presenter.ContentCreateUpdateRequest) error
	DeleteContent(id uint) error
	UpdateContentPosition(data presenter.UpdateContentPositionRequest) map[string]string
	GetContentDetails(id uint) (*presenter.ContentDetailResponse, error)

	AssignContentsToChapter(data models.ChapterContent) error
	//progress
	MarkAsCompleted(userID, contentID uint) error
	EvaluateProgress(data presenter.ProgressPresenter) error
	EvaluateContentProgress(data presenter.ContentProgressPresenter) error
	GetCourseIDByContentID(contentID uint) (uint, error)
	GetConsumedContentCount(userID, courseID uint) (int64, error)
	GetTotalContentCount(courseID uint) (int64, error)
	SaveOrUpdateProgress(progress *models.StudentCourse) error

	CheckStudentContent(userID, contentID uint) (models.StudentContent, error)
	GetContentProgressByContentID(contentID, userID uint) (models.StudentContent, error)
	AssignContentsToRelation(uint, uint, []uint) error
}
