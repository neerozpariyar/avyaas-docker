package usecase

/*
DeleteTest is a usecase method responsible for deleting a test with the specified ID.

Parameters:
  - id: The ID of the test to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (uCase *usecase) DeleteTest(id uint) error {
	// Checks if the test with the given ID exists
	if _, err := uCase.repo.GetTestByID(id); err != nil {
		return err
	}

	// Delegate the deletion of test and return error that is returned, if any
	return uCase.repo.DeleteTest(id)
}
