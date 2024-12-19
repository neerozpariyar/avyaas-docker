package repository

import "avyaas/internal/domain/models"

func (repo *Repository) GetTermsAndConditionByID(id uint) (*models.TermsAndCondition, error) {
	var termsCondition models.TermsAndCondition

	err := repo.db.Where("id = ?", id).First(&termsCondition).Error
	if err != nil {
		return nil, err
	}
	return &termsCondition, err
}
