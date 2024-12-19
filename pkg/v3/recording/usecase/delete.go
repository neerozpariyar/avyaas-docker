package usecase

func (uCase *usecase) DeleteRecording(id uint) error {
	if _, err := uCase.repo.GetRecordingByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteRecording(id)
}
