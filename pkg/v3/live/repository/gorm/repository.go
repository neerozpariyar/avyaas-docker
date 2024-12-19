package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	db          *gorm.DB
	accountRepo interfaces.AccountRepository
}

func New(db *gorm.DB, accountRepo interfaces.AccountRepository) *Repository {
	return &Repository{
		db:          db,
		accountRepo: accountRepo,
	}
}
