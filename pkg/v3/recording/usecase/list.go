package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) ListRecording(request presenter.RecordingListRequest) ([]models.Recording, int, error) {
	recordings, totalPage, err := u.repo.ListRecording(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	for i := range recordings {
		if recordings[i].Url != "" {
			recordings[i].Url = utils.GetFileURL(recordings[i].Url)
		}
	}

	return recordings, int(totalPage), nil
}
