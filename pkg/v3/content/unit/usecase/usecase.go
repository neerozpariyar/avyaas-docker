package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.UnitRepository
	courseRepo  interfaces.CourseRepository
	subjectRepo interfaces.SubjectRepository
	chapterRepo interfaces.ChapterRepository
}

func New(repo interfaces.UnitRepository, courseRepo interfaces.CourseRepository, subjectRepo interfaces.SubjectRepository, chapterRepo interfaces.ChapterRepository) *usecase {
	return &usecase{
		repo:        repo,
		courseRepo:  courseRepo,
		subjectRepo: subjectRepo,
		chapterRepo: chapterRepo,
	}
}
