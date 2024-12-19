package gorm

import (
	"avyaas/internal/domain/models"
)

/*
DeleteTestType is a repository method responsible for deleting a test type with the specified ID.

Parameters:
  - id: The ID of the test type to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (repo *Repository) DeleteTestType(id uint) error {
	// Perform a hard delete of the test type with the given ID using the GORM Unscoped method
	return repo.db.Unscoped().Where("id = ?", id).Delete(&models.TestType{}).Error
}
