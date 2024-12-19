package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdatePoll(poll models.Poll) map[string]string {
	var err error
	errMap := make(map[string]string)

	_, err = uCase.repo.GetPollByID(poll.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if err = uCase.repo.UpdatePoll(poll); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
