package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) UpdateNote(data presenter.NoteCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing note  with the provided note 's ID
	not, err := uCase.repo.GetNoteByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	notByID, err := uCase.repo.GetNoteByID(data.ID)
	if err == nil {
		// Check if the noteID is the same as of the requested note
		if not.ID != notByID.ID {
			errMap["noteID"] = fmt.Errorf("note  with  id: '%v' already exists", notByID.ID).Error()
			return errMap
		}
	}

	if _, err = uCase.contentRepo.GetContentByID(data.ContentID); err != nil {
		errMap["ContentID"] = err.Error()
		return errMap
	}

	// Delegate the update of note
	if err = uCase.repo.UpdateNote(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
