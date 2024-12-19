package repository

import "gorm.io/gorm"

/*
repository represents the authentication repository, which encapsulates the Gorm database connection
for handling authentication-related data operations.
*/
type Repository struct {
	db *gorm.DB
}

/*
New creates and returns a new instance of the authentication repository. It takes a Gorm database
connection as a parameter, which is used for data access.
*/
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
