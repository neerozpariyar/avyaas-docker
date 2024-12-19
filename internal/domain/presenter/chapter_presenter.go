package presenter

import "avyaas/internal/domain/models"

type ChapterCreateUpdateRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title" validate:"required"`
	// UnitID    uint   `json:"unitID" validate:"required"`
	Position  uint
	SubjectID uint `json:"subjectID"`
}

type UpdateChapterPositionRequest struct {
	ChapterIDs []uint `json:"chapterIDs"`
}

type ChapterListResponse struct {
	Success     bool  `json:"success"`
	CurrentPage int32 `json:"currentPage"`
	TotalPage   int32 `json:"totalPage"`
	// TotalCount  int32     `json:"totalCount"`
	Data []Chapter `json:"data"`
}

type ChapterListRequest struct {
	Page          int
	ChapterFilter FilterChapterRequest
	Search        string
	PageSize      int
}

type FilterChapterRequest struct {
	UnitID    uint `json:"unitID" validate:"required"`
	SubjectID uint `json:"subjectID" validate:"required"`
}

type Chapter struct {
	ID       uint             `json:"id"`
	Title    string           `json:"title"`
	Unit     interface{}      `json:"unit"`
	Position int              `json:"position"`
	Contents []models.Content `json:"chapterContents"`
}

type AssignContentsToRelation struct {
	SubjectID  uint   `json:"subjectID"`
	ContentIDs []uint `json:"contentIDs"`
	RelationID uint   `json:"relationID"`
}
