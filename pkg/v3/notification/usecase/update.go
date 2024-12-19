package usecase

import (
	"avyaas/internal/domain/models"
)

func (uCase *usecase) UpdateNotification(data models.Notification) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing feedback  with the provided feedback's ID
	_, err = uCase.repo.GetNotificationByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Delegate the update of feedback
	if err = uCase.repo.UpdateNotification(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
