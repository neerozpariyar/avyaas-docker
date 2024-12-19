package repository

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateFaq(data *models.FAQ) error {
	updateFaq := &models.FAQ{
		Title:       data.Title,
		Description: data.Description,
	}

	var err error

	if err = repo.db.Where("id = ?", data.ID).Updates(&updateFaq).Error; err != nil {
		return err
	}
	return err
}
