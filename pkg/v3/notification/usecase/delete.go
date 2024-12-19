package usecase

func (uCase *usecase) DeleteNotification(id uint) error {
	// Checks if the notification  with the given ID exists
	if _, err := uCase.repo.GetNotificationByID(id); err != nil {
		return err
	}

	// Delegate the deletion of notification
	return uCase.repo.DeleteNotification(id)
}
