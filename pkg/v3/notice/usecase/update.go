package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) UpdateNotice(requestBody presenter.NoticeCreateUpdatePresenter) (*models.Notice, map[string]string) {
	var err error

	errMap := make(map[string]string)

	noticeID, err := uCase.repo.GetNoticeByID(requestBody.ID)
	if err != nil {
		errMap["id"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateNotice(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return noticeID, errMap
}
