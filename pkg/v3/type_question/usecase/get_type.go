package usecase

func (u *usecase) GetQuestionTypeByID(id uint) (string, error) {
	typ, err := u.repo.GetTypeQuestionByID(id)
	if err != nil {
		return "", err
	}

	return typ.Type, err
}
