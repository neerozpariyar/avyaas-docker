package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	db         *gorm.DB
	courseRepo interfaces.CourseRepository
}

func New(db *gorm.DB, courseRepo interfaces.CourseRepository) *Repository {
	return &Repository{
		db:         db,
		courseRepo: courseRepo,
	}
}
