package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetPollByID(id uint) (models.Poll, error) {
	var poll models.Poll

	// Retrieve the poll from the database based on given id
	err := repo.db.Where("id = ?", id).First(&poll).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Poll{}, fmt.Errorf("poll with poll id: '%d' not found", id)
		}

		return models.Poll{}, err
	}

	return poll, nil
}

func (repo *Repository) GetPollOptionByID(id uint) (models.PollOption, error) {
	var pollOption models.PollOption

	// Retrieve the poll option from the database based on given id
	err := repo.db.Where("id = ?", id).First(&pollOption).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.PollOption{}, fmt.Errorf("poll option with id: '%d' not found", id)
		}

		return models.PollOption{}, err
	}

	return pollOption, nil
}

func (repo *Repository) GetVoteCountForOption(pollID uint, optionName string) (int, error) {
	var voteCount int

	err := repo.db.Debug().Model(&models.PollOption{}).
		Select("COUNT(poll_votes.id) as vote_count").
		Joins("LEFT JOIN poll_votes ON poll_options.id = poll_votes.poll_option_id").
		Where("poll_options.poll_id = ? AND poll_options.option = ?", pollID, optionName).
		Group("poll_options.id").
		Scan(&voteCount).Error

	if err != nil {
		return 0, err
	}

	return voteCount, nil
}

func (repo *Repository) GetVotedOptionByUserID(userID, pollID uint) (uint, error) {
	var votedOption models.PollVote
	//    SELECT poll_option_id FROM poll_votes WHERE (user_id = 25 AND poll_id = 2) AND poll_votes.deleted_at IS NULL ORDER BY poll_votes.id LIMIT 1;
	err := repo.db.Debug().Model(&models.PollVote{}).
		Select("poll_option_id").
		Where("user_id = ? AND poll_id = ?", userID, pollID).
		First(&votedOption).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}

	return votedOption.PollOptionID, nil
}
