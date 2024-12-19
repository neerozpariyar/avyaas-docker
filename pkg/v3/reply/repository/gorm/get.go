package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetReplyByID(id uint) (models.Reply, error) {
	var reply models.Reply

	// Retrieve the reply from the database based on given id
	err := repo.db.Where("id = ?", id).First(&reply).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Reply{}, fmt.Errorf("reply with reply id: '%d' not found", id)
		}

		return models.Reply{}, err
	}

	return reply, nil
}
