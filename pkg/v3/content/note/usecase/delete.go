package usecase

func (uCase *usecase) DeleteNote(id uint) error {
	if _, err := uCase.repo.GetNoteByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteNote(id)
}
