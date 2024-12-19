package gorm

import (
	"avyaas/internal/domain/models"
)

/*
UpdateTestType is a repository method responsible for updating the details of a test type in the
repository based on the provided test type's ID.

Parameters:
  - testType: A models.TestType instance containing the updated details of the test type.

Returns:
  - err: An error, if any, encountered during the update operation.
*/
func (repo *Repository) UpdateTestType(testType models.TestType) error {
	return repo.db.Debug().Model(&models.TestType{}).Where("id = ?", testType.ID).Update("title", testType.Title).Error
}
