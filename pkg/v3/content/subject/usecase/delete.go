package usecase

func (uCase *usecase) DeleteSubject(id uint) error {
	// Checks if the subject  with the given ID exists
	if _, err := uCase.repo.GetSubjectByID(id); err != nil {
		return err
	}

	// Delegate the deletion of subject
	return uCase.repo.DeleteSubject(id)
}
