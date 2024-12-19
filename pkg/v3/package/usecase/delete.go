package usecase

func (uCase *usecase) DeletePackage(id uint) error {
	// Checks if the unit  with the given ID exists
	if _, err := uCase.repo.GetPackageByID(id); err != nil {
		return err
	}

	// Delegate the deletion of unit
	return uCase.repo.DeletePackage(id)
}
