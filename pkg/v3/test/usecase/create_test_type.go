package usecase

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
CreateTestType is a usecase method responsible for creating a new test type.

Parameters:
  - testType: A models.TestType instance representing the test type to be created.

Returns:
  - errMap: A map[string]string containing error messages, if any, encountered during the process.
*/
func (uCase *usecase) CreateTestType(testType models.TestType) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Check if a test type with the provided title already exists
	_, err = uCase.repo.GetTestTypeByName(testType.Title)
	if err == nil {
		// If a test type with the given title already exists, return an error
		errMap["title"] = fmt.Errorf("test type with title: '%s' already exists", testType.Title).Error()
		return errMap
	}

	// If no test is found with the provided title, create the test type
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = uCase.repo.CreateTestType(testType); err != nil {
			errMap["error"] = err.Error()
			// return errMap
		}
	} else {
		errMap["error"] = err.Error()
	}

	return errMap

}
