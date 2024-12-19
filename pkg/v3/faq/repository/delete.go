package repository

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteFaq(id uint) (*models.FAQ, error) {
	var err error

	_, err = repo.GetFAQByID(id)
	if err != nil {
		return nil, err
	}

	err = repo.db.Where("id = ?", id).Delete(&models.FAQ{}).Error
	if err != nil {
		return nil, err
	}
	return nil, err
}
