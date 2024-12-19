package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"errors"
)

func (repo *Repository) CreatePoll(data presenter.PollCreateUpdateRequest) error {
	transaction := repo.db.Begin()
	if len(data.Options) != 4 {
		return errors.New("4 options needed")
	}

	poll := &models.Poll{
		Question:  data.Question,
		SubjectID: data.SubjectID,
		CourseID:  data.CourseID,
		UserID:    data.CreatedBy,
	}
	var pollOptions []models.PollOption
	for _, option := range data.Options {
		pollOption := models.PollOption{
			Option: option,
			UserID: data.CreatedBy,
		}
		pollOptions = append(pollOptions, pollOption)
	}

	poll.Options = pollOptions

	if err := transaction.Create(poll).Error; err != nil {

		transaction.Rollback()
		return err

	}

	transaction.Commit()
	return nil

}
