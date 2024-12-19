package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) CreateFeedback(data models.Feedback) map[string]string {
	var err error
	errMap := make(map[string]string)

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}
	
	if err = uCase.repo.CreateFeedback(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
