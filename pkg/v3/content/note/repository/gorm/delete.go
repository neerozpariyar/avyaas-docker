package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteNote(id uint) error {
	transaction := repo.db.Begin()

	note, err := repo.GetNoteByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if note.File != "" {
		var nFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", note.File).First(nFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", nFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	// Perform a hard delete of the note  with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Note{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
