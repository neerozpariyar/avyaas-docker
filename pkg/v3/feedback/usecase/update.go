package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdateFeedback(data models.Feedback) map[string]string {
	var err error
	errMap := make(map[string]string)

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	// Retrieve the existing feedback  with the provided feedback's ID
	_, err = uCase.repo.GetFeedbackByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Delegate the update of feedback
	if err = uCase.repo.UpdateFeedback(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
