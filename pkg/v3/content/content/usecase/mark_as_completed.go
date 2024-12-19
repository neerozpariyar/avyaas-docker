package usecase

func (uCase *usecase) MarkAsCompleted(userID, contentID uint) map[string]string {
	var err error

	errMap := make(map[string]string)

	if _, err := uCase.repo.GetContentByID(contentID); err != nil {
		errMap["contentID"] = err.Error()
		return errMap
	}

	if err = uCase.repo.MarkAsCompleted(userID, contentID); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
