package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteRecording(id uint) error {
	transaction := repo.db.Begin()

	recording, err := repo.GetRecordingByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if recording.Url != "" {
		var cFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", recording.Url).First(&cFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", cFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	// Perform a hard delete of the recording  with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Recording{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
