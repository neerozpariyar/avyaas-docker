package usecase

/*
DeleteQuestion is a usecase method responsible for deleting a question with the specified ID.

Parameters:
  - id: The ID of the question to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (uCase *usecase) DeleteTypeQuestion(id uint) error {
	// Checks if the question with the given ID exists
	if _, err := uCase.repo.GetTypeQuestionByID(id); err != nil {
		return err
	}

	// Delegate the deletion of question
	return uCase.repo.DeleteTypeQuestion(id)
}
