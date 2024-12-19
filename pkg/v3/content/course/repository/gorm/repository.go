package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	db              *gorm.DB
	accountRepo     interfaces.AccountRepository
	courseGroupRepo interfaces.CourseGroupRepository
}

func New(db *gorm.DB, accountRepo interfaces.AccountRepository, courseGroupRepo interfaces.CourseGroupRepository) *Repository {
	return &Repository{
		db:              db,
		accountRepo:     accountRepo,
		courseGroupRepo: courseGroupRepo,
	}
}
