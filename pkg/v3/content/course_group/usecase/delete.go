package usecase

/*
DeleteCourseGroup is a usecase method responsible for deleting a course group with the specified ID.

Parameters:
  - id: The ID of the course group to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (uCase *usecase) DeleteCourseGroup(id uint) error {
	// Checks if the course group with the given ID exists
	if _, err := uCase.repo.GetCourseGroupByID(id); err != nil {
		return err
	}

	// Delegate the deletion of course group
	return uCase.repo.DeleteCourseGroup(id)
}
