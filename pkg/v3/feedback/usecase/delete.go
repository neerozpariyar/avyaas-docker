package usecase

func (uCase *usecase) DeleteFeedback(id uint) error {
	// Checks if the feedback  with the given ID exists
	if _, err := uCase.repo.GetFeedbackByID(id); err != nil {
		return err
	}

	// Delegate the deletion of feedback
	return uCase.repo.DeleteFeedback(id)
}
