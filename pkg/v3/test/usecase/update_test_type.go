package usecase

import (
	"avyaas/internal/domain/models"
	"fmt"
)

/*
UpdateTestType is a usecase method for updating the test type in the repository.

Parameters:
  - testType: A models.TestType struct containing the updated details of the test type.

Returns:
  - errMap: A map containing error messages, if any, encountered during the update operation.
*/
func (uCase *usecase) UpdateTestType(testType models.TestType) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing test type with the provided test type's ID
	testTypeByID, err := uCase.repo.GetTestTypeByID(testType.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Check if a test type with the given title already exists
	testTypeByName, err := uCase.repo.GetTestTypeByName(testType.Title)
	if err == nil {
		// Check if the title is the same as of the requested test type
		if testTypeByID.Title != testTypeByName.Title {
			errMap["title"] = fmt.Errorf("test type with title: '%s' already exists", testType.Title).Error()
			return errMap
		}
	}

	// Delegate the update of test type
	if err = uCase.repo.UpdateTestType(testType); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
