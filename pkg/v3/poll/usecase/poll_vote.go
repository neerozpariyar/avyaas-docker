package usecase

import "avyaas/internal/domain/presenter"

func (uCase *usecase) PollVote(request presenter.PollVoteRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	_, err = uCase.repo.GetPollByID(request.PollID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	_, err = uCase.repo.GetPollOptionByID(request.OptionID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if err = uCase.repo.PollVote(request); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
