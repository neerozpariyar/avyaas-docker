package usecase

/*
DeleteTestType is a usecase method responsible for deleting a test type with the specified ID.

Parameters:
  - id: The ID of the test type to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (uCase *usecase) DeleteTestType(id uint) error {
	// Checks if the test type with the given ID exists
	if _, err := uCase.repo.GetTestTypeByID(id); err != nil {
		return err
	}

	// Delegate the deletion of test type
	return uCase.repo.DeleteTestType(id)
}
