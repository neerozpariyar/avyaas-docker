package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

/*
repository represents the authentication repository, which encapsulates the Gorm database connection
for handling authentication-related data operations.
*/
type Repository struct {
	db          *gorm.DB
	accountRepo interfaces.AccountRepository
	courseRepo  interfaces.CourseRepository
}

/*
New creates and returns a new instance of the authentication repository. It takes a Gorm database
connection as a parameter, which is used for data access.
*/
func New(db *gorm.DB, accountRepo interfaces.AccountRepository, courseRepo interfaces.CourseRepository) *Repository {
	return &Repository{
		db:          db,
		accountRepo: accountRepo,
		courseRepo:  courseRepo,
	}
}
