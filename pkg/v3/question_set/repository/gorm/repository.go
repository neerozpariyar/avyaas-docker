package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

/*
repository represents the question set repository, which encapsulates the Gorm database connection
for handling question set related data operations.
*/
type Repository struct {
	db           *gorm.DB
	questionRepo interfaces.QuestionRepository
}

/*
New creates and returns a new instance of the question set repository. It takes a Gorm database
connection as a parameter, which is used for data access.
*/
func New(db *gorm.DB, questionRepo interfaces.QuestionRepository) *Repository {
	return &Repository{
		db:           db,
		questionRepo: questionRepo,
	}
}
