package usecase

func (uCase *usecase) DeleteLiveGroup(id uint) error {
	if _, err := uCase.repo.GetLiveGroupByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteLiveGroup(id)
}
