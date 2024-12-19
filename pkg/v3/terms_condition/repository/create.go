package repository

import "avyaas/internal/domain/models"

func (repo *Repository) CreateTermsAndCondition(data *models.TermsAndCondition) error {
	transaction := repo.db.Begin()

	termsAndCondition := &models.TermsAndCondition{
		Title:       data.Title,
		Description: data.Description,
	}

	if err := transaction.Create(&termsAndCondition).Error; err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
