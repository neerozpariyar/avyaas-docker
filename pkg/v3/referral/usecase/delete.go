package usecase

func (uCase *usecase) DeleteReferral(id uint) error {
	if _, err := uCase.repo.GetReferralByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteReferral(id)
}
