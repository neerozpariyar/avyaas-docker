package usecase

func (u *usecase) GetPackageTypeServices(packageTypeID uint) ([]uint, error) {
	serviceIDs, err := u.repo.GetPackageTypeServices(packageTypeID)
	if err != nil {
		return nil, err
	}

	return serviceIDs, nil
}
