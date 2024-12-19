package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) GetTotalContentCount(courseID uint) (int64, error) {
	var count int64

	if courseID != 0 {
		var contentIDs []uint
		query := repo.db.Model(&models.Content{}).
			Joins("JOIN chapter_contents ON contents.id = chapter_contents.content_id").
			Joins("JOIN chapters ON chapter_contents.chapter_id = chapters.id").
			Joins("JOIN units ON chapters.unit_id = units.id").
			Joins("JOIN subjects ON units.subject_id = subjects.id").
			Where("subjects.course_id = ?", courseID).
			Distinct("contents.id").
			Count(&count).
			Pluck("contents.id", &contentIDs)
		if query.Error != nil {
			fmt.Printf("Error executing query: %v\n", query.Error)
			return 0, query.Error
		}

	}

	return count, nil
}
