package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) GetCourseIDByContentID(contentID uint) (uint, error) {
	var studentContent models.StudentContent

	if err := repo.db.Where("content_id = ?", contentID).First(&studentContent).Error; err != nil {
		return 0, err
	}

	return studentContent.CourseID, nil
}
