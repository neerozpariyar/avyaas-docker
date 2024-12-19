package usecase

func (uCase *usecase) DeletePackageType(id uint) error {
	// Checks if the package type with the given ID exists
	if _, err := uCase.repo.GetPackageTypeByID(id); err != nil {
		return err
	}

	// Delegate the deletion of package type
	return uCase.repo.DeletePackageType(id)
}
