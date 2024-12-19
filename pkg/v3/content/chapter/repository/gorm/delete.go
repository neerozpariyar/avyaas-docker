package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteChapter(id uint) error {
	var chapter models.Chapter

	err := repo.db.Where("id  = ?", id).First(&chapter).Error

	if err != nil {
		return err

	}
	// Perform a hard delete of the chapter  with the given ID using the GORM Unscoped method
	return repo.db.Unscoped().Where("id = ?", id).Delete(&chapter).Error
}
