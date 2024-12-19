package usecase

func (uCase *usecase) DeleteComment(id uint) error {
	if _, err := uCase.repo.GetCommentByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteComment(id)
}
