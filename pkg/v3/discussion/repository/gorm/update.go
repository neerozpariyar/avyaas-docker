package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) UpdateDiscussion(discussion presenter.DiscussionCreateUpdateRequest) error {
	return repo.db.Where("id = ?", discussion.ID).Updates(&models.Discussion{
		Title:     discussion.Title,
		Query:     discussion.Query,
		SubjectID: discussion.SubjectID,
		CourseID:  discussion.CourseID,
	}).Error
}
