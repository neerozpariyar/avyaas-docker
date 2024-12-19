package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) CreateNote(data presenter.NoteCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	if _, err = uCase.contentRepo.GetContentByID(data.ContentID); err != nil {
		errMap["ContentID"] = err.Error()
		return errMap
	}

	if err = uCase.repo.CreateNote(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
