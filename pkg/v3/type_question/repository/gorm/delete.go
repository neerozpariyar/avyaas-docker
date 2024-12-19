package gorm

import (
	"avyaas/internal/domain/models"
)

/*
DeleteQuestion is a repository method responsible for deleting a question with the specified ID.

Parameters:
  - id: The ID of the question to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (repo *Repository) DeleteTypeQuestion(id uint) error {
	transaction := repo.db.Begin()

	question, err := repo.GetTypeQuestionByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	options, _ := repo.GetTypeOptionsByQuestionID(question.ID)
	for _, option := range options {
		if option.Image != nil && *option.Image != "" {
			var imageFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", option.Image).First(&imageFile).Error
			if err == nil {
				if err = transaction.Model(models.File{}).Where("id = ?", imageFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		if option.Audio != nil && *option.Audio != "" {
			var audioFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", option.Audio).First(&audioFile).Error
			if err == nil {
				if err = transaction.Model(models.File{}).Where("id = ?", audioFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}
	}

	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.TypeQuestion{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
