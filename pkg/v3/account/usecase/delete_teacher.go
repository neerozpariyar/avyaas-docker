package usecase

/*
DeleteTeacher is a use case function responsible for deleting a teacher.

Parameters:
  - uCase: A pointer to the use case struct, representing the business logic for user-related
    operations. It is used to access the repository for deleting a teacher.
  - id: An unsigned integer representing the ID of the teacher to be deleted.

Returns:
  - error: An error indicating any issues encountered during the deletion of the teacher.
    A nil error signifies a successful deletion.
*/
func (uCase *usecase) DeleteTeacher(id uint) error {
	// Checks if the user with the given ID exists
	if _, err := uCase.repo.GetUserByID(id); err != nil {
		return err
	}

	// Delegate the deletion of teacher
	return uCase.repo.DeleteTeacher(id)
}
