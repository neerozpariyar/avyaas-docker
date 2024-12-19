package usecase

func (uCase *usecase) DeleteContent(id uint) error {
	if _, err := uCase.repo.GetContentByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteContent(id)
}
