package repository

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteTermsAndCondition(id uint) (*models.TermsAndCondition, error) {
	var err error

	_, err = repo.GetTermsAndConditionByID(id)
	if err != nil {
		return nil, err
	}

	err = repo.db.Where("id = ?", id).Delete(&models.TermsAndCondition{}).Error
	if err != nil {
		return nil, err
	}
	return nil, err
}
