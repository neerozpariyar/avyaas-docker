package repository

import "avyaas/internal/domain/models"

func (repo *Repository) CreateFaq(data *models.FAQ) error {
	transaction := repo.db.Begin()

	faq := &models.FAQ{
		Title:       data.Title,
		Description: data.Description,
	}

	if err := transaction.Create(&faq).Error; err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
