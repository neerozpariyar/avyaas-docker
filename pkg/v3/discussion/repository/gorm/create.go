package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) CreateDiscussion(data presenter.DiscussionCreateUpdateRequest) error {
	transaction := repo.db.Begin()
	err := transaction.Create(&models.Discussion{
		Title:     data.Title,
		Query:     data.Query,
		SubjectID: data.SubjectID,
		CourseID:  data.CourseID,
		UserID:    data.CreatedBy,
	}).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
