package repository

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateTermsAndCondition(data *models.TermsAndCondition) error {
	var err error

	updateTermsAndCondition := &models.TermsAndCondition{
		Title:       data.Title,
		Description: data.Description,
	}

	if err = repo.db.Where("id = ?", data.ID).Updates(&updateTermsAndCondition).Error; err != nil {
		return err
	}
	return err
}
