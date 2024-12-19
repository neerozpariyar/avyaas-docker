package repository

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteNotice(id uint) (*models.Notice, error) {
	var err error

	err = repo.db.Where("id = ?", id).Delete(&models.Notice{}).Error
	if err != nil {
		return nil, err
	}

	return nil, err
}
