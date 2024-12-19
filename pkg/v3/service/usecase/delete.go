package usecase

func (uCase *usecase) DeleteService(id uint) error {
	// Checks if the service with the given ID exists
	if _, err := uCase.repo.GetServiceByID(id); err != nil {
		return err
	}

	// Delegate the deletion of service
	return uCase.repo.DeleteService(id)
}
