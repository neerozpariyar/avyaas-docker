package usecase

/*
DeleteQuestionSet is a usecase method responsible for deleting a question set with the specified ID.

Parameters:
  - id: The ID of question set to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (uCase *usecase) DeleteQuestionSet(id uint) error {
	// Checks if the question set with the given ID exists
	if _, err := uCase.repo.GetQuestionSetByID(id); err != nil {
		return err
	}

	// Delegate the deletion of question set and return error that is returned, if any
	return uCase.repo.DeleteQuestionSet(id)
}
