package usecase

import "avyaas/internal/domain/presenter"

func (usecase *usecase) GetSubjectHeirarchy(id uint, userID uint) ([]presenter.SubjectHeirarchyDetails, error) {
	if _, err := usecase.repo.GetSubjectByID(id); err != nil {
		return nil, err
	}

	if _, err := usecase.accountRepo.GetUserByID(userID); err != nil {
		return nil, err
	}

	result, err := usecase.repo.GetSubjectHeirarchy(id, userID)

	if err != nil {
		return nil, err
	}

	return result, nil
}
