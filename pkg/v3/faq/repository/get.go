package repository

import "avyaas/internal/domain/models"

func (repo *Repository) GetFAQByID(id uint) (*models.FAQ, error) {
	var faq models.FAQ

	err := repo.db.Where("id = ? ", id).First(&faq).Error
	if err != nil {
		return nil, err
	}
	return &faq, err
}
