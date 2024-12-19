package usecase

func (uCase *usecase) DeleteUnit(id uint) error {
	// Checks if the unit  with the given ID exists
	if _, err := uCase.repo.GetUnitByID(id); err != nil {
		return err
	}

	// Delegate the deletion of unit
	return uCase.repo.DeleteUnit(id)
}
