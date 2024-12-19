package usecase

func (uCase *usecase) DeleteReply(id uint) error {
	if _, err := uCase.repo.GetReplyByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteReply(id)
}
