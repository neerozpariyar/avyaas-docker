package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type PollUsecase interface {
	CreatePoll(data presenter.PollCreateUpdateRequest) map[string]string
	ListPoll(request presenter.PollListRequest) ([]presenter.Poll, int, error)
	UpdatePoll(poll models.Poll) map[string]string
	DeletePoll(id uint) error
	PollVote(request presenter.PollVoteRequest) map[string]string
}

type PollRepository interface {
	GetPollByID(id uint) (models.Poll, error)
	GetPollOptionByID(id uint) (models.PollOption, error)

	CreatePoll(data presenter.PollCreateUpdateRequest) error
	ListPoll(request presenter.PollListRequest) ([]models.Poll, float64, error)
	// GetTotalVoteCount(pollID uint) (map[string]int, error)
	GetVoteCountForOption(pollID uint, optionName string) (int, error)
	GetVotedOptionByUserID(userID, pollID uint) (uint, error)
	UpdatePoll(poll models.Poll) error
	DeletePoll(id uint) error
	PollVote(request presenter.PollVoteRequest) error
}
