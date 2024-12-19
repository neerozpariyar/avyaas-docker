package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) CreateReply(data presenter.ReplyCreateUpdateRequest) error {
	transaction := repo.db.Begin()

	err := transaction.Create(&models.Reply{
		Reply:        data.Reply,
		CourseID:     data.CourseID,
		DiscussionID: data.DiscussionID,
		UserID:       data.CreatedBy,
	}).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return err
}
	