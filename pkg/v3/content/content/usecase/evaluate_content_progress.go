package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) EvaluateContentProgress(data presenter.ContentProgressPresenter) error {
	courseID, err := uCase.repo.GetCourseIDByContentID(data.ID)
	if err != nil {
		return err
	}

	// Invoke usecase to evaluate overall progress
	progressRequest := presenter.ProgressPresenter{
		UserID:   data.UserID,
		CourseID: courseID,
	}
	err = uCase.EvaluateProgress(progressRequest)
	if err != nil {
		return err
	}
	return uCase.repo.EvaluateContentProgress(data)
}
