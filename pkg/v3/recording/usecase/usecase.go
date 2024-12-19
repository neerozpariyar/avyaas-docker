package usecase

import (
	"avyaas/internal/domain/interfaces"
)

type usecase struct {
	repo     interfaces.RecordingRepository
	liveRepo interfaces.LiveRepository
}

func New(repo interfaces.RecordingRepository, liveRepo interfaces.LiveRepository) *usecase {
	return &usecase{
		repo:     repo,
		liveRepo: liveRepo,
	}
}
