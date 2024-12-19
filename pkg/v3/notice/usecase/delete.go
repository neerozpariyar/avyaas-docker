package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) DeleteNotice(id uint) (*models.Notice, error) {
	var err error

	_, err = uCase.repo.GetNoticeByID(id)
	if err != nil {
		return nil, err
	}

	notice, err := uCase.repo.DeleteNotice(id)
	if err != nil {
		return nil, err
	}

	return notice, err
}
