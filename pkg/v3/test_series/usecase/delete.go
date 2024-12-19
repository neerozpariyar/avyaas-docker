package usecase

func (uCase *usecase) DeleteTestSeries(id uint) error {
	// Checks if the test series with the given ID exists
	if _, err := uCase.repo.GetTestSeriesByID(id); err != nil {
		return err
	}

	// Delegate the deletion of test series
	return uCase.repo.DeleteTestSeries(id)
}
