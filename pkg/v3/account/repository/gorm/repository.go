package gorm

import (
	"gorm.io/gorm"
)

/*
repository represents the account repository, which encapsulates the Gorm database connection for
handling account related data operations.
*/
type Repository struct {
	db *gorm.DB
}

/*
New creates and returns a new instance of the account repository. It takes a Gorm database connection
as a parameter, which is used for data access.
*/
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
