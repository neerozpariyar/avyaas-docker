package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) ListNote(data presenter.NoteListRequest) ([]models.Note, int, error) {
	notes, totalPage, err := u.repo.ListNote(data)
	if err != nil {
		return nil, int(totalPage), err
	}

	for i := range notes {
		if notes[i].File != "" {
			notes[i].File = utils.GetFileURL(notes[i].File)
		}
	}

	return notes, int(totalPage), nil
}
