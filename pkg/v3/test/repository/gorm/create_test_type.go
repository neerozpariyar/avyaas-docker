package gorm

import (
	"avyaas/internal/domain/models"
)

/*
CreateTestType is a repository method responsible for creating a new test type in the database.

Parameters:
  - testType: A models.TestType instance representing the test type to be created in the database.

Returns:
  - error: An error, if any, encountered during the database insertion operation.
*/
func (repo *Repository) CreateTestType(testType models.TestType) error {
	return repo.db.Create(&testType).Error
}
