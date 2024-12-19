package repository

import "avyaas/internal/domain/models"

func (repo *Repository) GetNoticeByID(id uint) (*models.Notice, error) {
	var notice models.Notice

	err := repo.db.Where("id = ?", id).First(&notice).Error
	if err != nil {
		return nil, err
	}

	return &notice, err
}
