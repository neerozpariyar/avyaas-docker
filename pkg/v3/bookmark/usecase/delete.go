package usecase

func (uCase *usecase) DeleteBookmark(id uint) error {
	// Checks if the bookmark  with the given ID exists
	if _, err := uCase.repo.GetBookmarkByID(id); err != nil {
		return err
	}

	// Delegate the deletion of bookmark
	return uCase.repo.DeleteBookmark(id)
}
