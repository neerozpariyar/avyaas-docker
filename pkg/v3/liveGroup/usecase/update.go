package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdateLiveGroup(data models.LiveGroup) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing liveGroup  with the provided liveGroup 's ID
	_, err = uCase.repo.GetLiveGroupByID(data.ID)
	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseGroupID"] = err.Error()
		return errMap
	}

	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Delegate the update of liveGroup
	if err = uCase.repo.UpdateLiveGroup(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
