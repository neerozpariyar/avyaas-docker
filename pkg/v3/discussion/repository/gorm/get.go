package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetDiscussionByID(id uint) (models.Discussion, error) {
	var discussion models.Discussion

	// Retrieve the discussion from the database based on given id
	err := repo.db.Where("id = ?", id).First(&discussion).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Discussion{}, fmt.Errorf("discussion with discussion id: '%d' not found", id)
		}

		return models.Discussion{}, err
	}

	return discussion, nil
}
func (repo *Repository) GetHasLikedValue(discussionID, userID uint) (bool, error) {
	var vote models.Vote
	err := repo.db.Debug().Model(&models.Vote{}).
		Select("has_liked").
		Where("discussion_id = ? AND user_id = ?", discussionID, userID).
		First(&vote).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return vote.HasLiked, nil
}
