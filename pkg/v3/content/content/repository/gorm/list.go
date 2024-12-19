package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"math"
)

func (repo *Repository) ListContent(request presenter.ContentListRequest) ([]models.Content, float64, error) {
	var contents []models.Content
	baseQuery := repo.db.Debug().Model(&models.Content{}).Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id")

	if request.ContentFilter.ChapterID != 0 {

		baseQuery = baseQuery.Joins("JOIN unit_chapter_contents ucc ON contents.id = ucc.content_id").
			Joins("LEFT JOIN subject_unit_chapter_contents succ ON ucc.id = succ.unit_chapter_content_id").
			Where("ucc.unit_id = ? AND ucc.chapter_id = ? AND succ.subject_id = ?", request.ContentFilter.UnitID, request.ContentFilter.ChapterID, request.ContentFilter.SubjectID)

	}
	if request.Search != "" {
		baseQuery = baseQuery.Where("title LIKE ?", "%"+request.Search+"%")
	}

	err := baseQuery.Find(&contents).Error

	if err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(len(contents)) / float64(request.PageSize))

	return contents, totalPage, nil
}
