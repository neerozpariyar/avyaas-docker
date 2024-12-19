package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) UpdateCourse(data presenter.CourseCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing course  with the provided course's ID
	c, err := uCase.repo.GetCourseByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	courseByID, err := uCase.repo.GetCourseByCourseID(data.CourseID)
	if err == nil {
		if c.CourseID != courseByID.CourseID {
			errMap["courseID"] = fmt.Errorf("course with course id: '%s' already exists", courseByID.CourseID).Error()
			return errMap
		}
	}

	for _, CourseGroupID := range data.CourseGroupIDs {

		if _, err := uCase.courseGroupRepo.GetCourseGroupByID(CourseGroupID); err != nil {
			errMap["courseGroupID"] = err.Error()
			return errMap
		}
	}

	// Delegate the update of course
	if err = uCase.repo.UpdateCourse(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
