package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) UploadRecording(data presenter.RecordingCreateUpdateRequest) map[string]string {
	var err error

	errMap := make(map[string]string)
	if _, err = uCase.liveRepo.GetLiveByID(data.LiveID); err != nil {
		errMap["live_id_error"] = err.Error()
		return errMap
	}

	if err = uCase.repo.UploadRecording(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
