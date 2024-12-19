package usecase

func (uCase *usecase) DeleteChapter(id uint) error {
	// Checks if the unit  with the given ID exists
	if _, err := uCase.repo.GetChapterByID(id); err != nil {
		return err
	}

	// Delegate the deletion of unit
	return uCase.repo.DeleteChapter(id)
}
