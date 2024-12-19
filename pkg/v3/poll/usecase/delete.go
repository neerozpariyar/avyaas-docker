package usecase

func (uCase *usecase) DeletePoll(id uint) error {
	if _, err := uCase.repo.GetPollByID(id); err != nil {
		return err
	}

	return uCase.repo.DeletePoll(id)
}
