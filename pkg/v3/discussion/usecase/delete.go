package usecase

func (uCase *usecase) DeleteDiscussion(id uint) error {

	if _, err := uCase.repo.GetDiscussionByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteDiscussion(id)
}
