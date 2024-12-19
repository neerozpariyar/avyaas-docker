package usecase

func (uCase *usecase) UpdateTestStatus(id uint) map[string]string {
	errMap := make(map[string]string)

	// Check if the test with the given ID exists
	test, err := uCase.repo.GetTestByID(id)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Call the repository to update the test status
	if err = uCase.repo.UpdateTestStatus(test); err != nil {
		return errMap
	}

	return errMap
}
