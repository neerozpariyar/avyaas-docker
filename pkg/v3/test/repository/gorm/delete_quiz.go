package gorm

import (
	"avyaas/internal/domain/models"

	"gorm.io/gorm"
)

/*
DeleteTestType is a repository method responsible for deleting a test type with the specified ID.

Parameters:
  - id: The ID of the test type to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (repo *Repository) DeleteTest(id uint) error {
	transaction := repo.db.Begin()

	err := repo.DeleteTestQuestionSetByTestID(id, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the test type with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Test{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}

func (repo *Repository) DeleteTestQuestionSetByTestID(testID uint, transaction *gorm.DB) error {
	return transaction.Where("test_id = ?", testID).Delete(&models.TestQuestionSet{}).Error
}
