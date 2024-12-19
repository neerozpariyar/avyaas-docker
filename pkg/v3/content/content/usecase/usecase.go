package usecase

import (
	"avyaas/internal/domain/interfaces"
)

type usecase struct {
	repo         interfaces.ContentRepository
	courseRepo   interfaces.CourseRepository
	chapterRepo  interfaces.ChapterRepository
	accountRepo  interfaces.AccountRepository
	bookmarkRepo interfaces.BookmarkRepository
	subjectRepo  interfaces.SubjectRepository
}

func New(repo interfaces.ContentRepository, courseRepo interfaces.CourseRepository, chapterRepo interfaces.ChapterRepository, accountRepo interfaces.AccountRepository,
	bookmarkRepo interfaces.BookmarkRepository, subjectRepo interfaces.SubjectRepository) *usecase {
	return &usecase{
		repo:         repo,
		courseRepo:   courseRepo,
		chapterRepo:  chapterRepo,
		accountRepo:  accountRepo,
		bookmarkRepo: bookmarkRepo,
		subjectRepo:  subjectRepo,
	}
}
