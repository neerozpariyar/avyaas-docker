package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdateReply(reply models.Reply) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing content  with the provided content 's ID
	existingReply, err := uCase.repo.GetReplyByID(reply.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if existingReply.UserID != reply.UserID {
		errMap["error"] = " not allowed to update others reply"
		return errMap

	}

	// Delegate the update of content
	if err = uCase.repo.UpdateReply(reply); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
