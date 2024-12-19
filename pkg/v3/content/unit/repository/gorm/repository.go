package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	db          *gorm.DB
	subjectRepo interfaces.SubjectRepository
}

func New(db *gorm.DB, subjectRepo interfaces.SubjectRepository) *Repository {
	return &Repository{
		db:          db,
		subjectRepo: subjectRepo,
	}
}
