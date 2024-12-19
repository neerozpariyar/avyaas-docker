package gorm

import (
	"avyaas/internal/domain/models"

	"gorm.io/gorm"
)

/*
DeleteQuestionSet is a repository method responsible for deleting a question set with the specified
ID.

Parameters:
  - id: The ID of the question set to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (repo *Repository) DeleteQuestionSet(id uint) error {
	transaction := repo.db.Begin()

	questionSet, err := repo.GetQuestionSetByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if questionSet.File != "" {
		var qsFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", questionSet.File).First(&qsFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", qsFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	err = repo.DeleteTestQuestionSetByQuestionSetID(id, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the question set with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.QuestionSet{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}

func (repo *Repository) DeleteTestQuestionSetByQuestionSetID(questionSetID uint, transaction *gorm.DB) error {
	return transaction.Where("question_set_id = ?", questionSetID).Delete(&models.TestQuestionSet{}).Error
}
