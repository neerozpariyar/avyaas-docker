package usecase

func (uCase *usecase) DeleteCourse(id uint) error {
	// Checks if the course with the given ID exists
	if _, err := uCase.repo.GetCourseByID(id); err != nil {
		return err
	}

	// Delegate the deletion of course
	return uCase.repo.DeleteCourse(id)
}
