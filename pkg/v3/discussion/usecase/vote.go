package usecase

func (uCase *usecase) LikeOrUnlikeDiscussion(discussionID, userID uint) map[string]string {
	var err error
	errMap := make(map[string]string)

	_, err = uCase.repo.GetDiscussionByID(discussionID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if err = uCase.repo.LikeOrUnlikeDiscussion(discussionID, userID); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
