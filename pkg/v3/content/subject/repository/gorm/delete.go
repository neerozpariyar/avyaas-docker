package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteSubject(id uint) error {
	transaction := repo.db.Begin()

	subject, err := repo.GetSubjectByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if subject.Thumbnail != "" {
		var sFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", subject.Thumbnail).First(&sFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", sFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	// Perform a hard delete of the subject group with the given ID using the GORM Unscoped method
	err = repo.db.Unscoped().Where("id = ?", id).Delete(&subject).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
