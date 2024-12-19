package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) DeletePoll(id uint) error {
	return repo.db.Unscoped().Where("id = ?", id).Delete(&models.Poll{}).Error

}

// transaction := repo.db.Begin()

// 	if err := transaction.Unscoped().Where("poll_id = ?", id).Delete(&models.PollVote{}).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	if err := transaction.Unscoped().Where("poll_id = ?", id).Delete(&models.PollOption{}).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	if err := transaction.Unscoped().Where("id = ?", id).Delete(&models.Poll{}).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	transaction.Commit()

// 	return nil
