package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteContent(id uint) error {
	transaction := repo.db.Begin()

	content, err := repo.GetContentByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if content.Url != "" {
		var cFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", content.Url).First(&cFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", cFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	err = transaction.Exec("DELETE FROM chapter_contents WHERE content_id = ?", id).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	// Perform a hard delete of the content  with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Content{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
