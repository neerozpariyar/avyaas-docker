package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

/*
repository represents the payment repository, which encapsulates the Gorm database connection for
handling payment related data operations.
*/
type Repository struct {
	db          *gorm.DB
	accountRepo interfaces.AccountRepository
}

/*
New creates and returns a new instance of the payment repository. It takes a Gorm database connection
as a parameter, which is used for data access.
*/
func New(db *gorm.DB, accountRepo interfaces.AccountRepository) *Repository {
	return &Repository{
		db:          db,
		accountRepo: accountRepo,
	}
}
