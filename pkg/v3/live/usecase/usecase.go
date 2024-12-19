package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo          interfaces.LiveRepository
	accountRepo   interfaces.AccountRepository
	liveGroupRepo interfaces.LiveGroupRepository
	courseRepo    interfaces.CourseRepository
	subjectRepo   interfaces.SubjectRepository
}

func New(repo interfaces.LiveRepository, accountRepo interfaces.AccountRepository, liveGroupRepo interfaces.LiveGroupRepository, courseRepo interfaces.CourseRepository, subjectRepo interfaces.SubjectRepository) *usecase {
	return &usecase{
		repo:          repo,
		accountRepo:   accountRepo,
		liveGroupRepo: liveGroupRepo,
		courseRepo:    courseRepo,
		subjectRepo:   subjectRepo,
	}
}
