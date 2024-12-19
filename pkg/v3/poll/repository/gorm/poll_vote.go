package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"errors"
)

func (repo *Repository) PollVote(request presenter.PollVoteRequest) error {
	// Check if the poll exists
	var existingPoll models.Poll
	if err := repo.db.First(&existingPoll, request.PollID).Error; err != nil {
		return errors.New("poll not found with the given ID")
	}

	transaction := repo.db.Begin()

	// Check if the user has already voted for this poll
	var existingVote models.PollVote
	if err := transaction.Where("user_id = ? AND poll_id = ?", request.UserID, request.PollID).First(&existingVote).Error; err == nil {
		return errors.New("already voted for this poll")
	}

	// Check if the pollOptionID exists
	var existingOption models.PollOption
	if err := repo.db.First(&existingOption, request.OptionID).Error; err != nil {
		return errors.New("poll option not found with the given ID")
	}

	if err := transaction.Create(&models.PollVote{
		UserID:       request.UserID,
		PollID:       request.PollID,
		PollOptionID: request.OptionID,
	}).Error; err != nil {
		return err
	}

	transaction.Commit()
	return nil
}
