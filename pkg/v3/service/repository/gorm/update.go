package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) UpdateService(data presenter.ServiceCreateUpdateRequest) error {
	// var oldUrlIDs []uint
	transaction := repo.db.Begin()

	// err := repo.db.Model(&models.ServiceUrl{}).Select("url_id").Where("service_id = ?", data.ID).Scan(&oldUrlIDs).Error
	// if err != nil {
	// 	return err
	// }

	// addUrls, delUrls := utils.CompareDifferences(oldUrlIDs, data.UrlIDs)
	// if len(addUrls) > 0 {
	// 	err = repo.CreateServiceUrls(data.ID, addUrls, transaction)
	// 	if err != nil {
	// 		transaction.Rollback()
	// 		return err
	// 	}
	// }

	// if len(delUrls) > 0 {
	// 	err = repo.DeleteServiceUrls(data.ID, delUrls, transaction)
	// 	if err != nil {
	// 		transaction.Rollback()
	// 		return err
	// 	}
	// }

	err := transaction.Where("id = ?", data.ID).Updates(&models.Service{
		Title:       data.Title,
		Description: data.Description,
	}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
