package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
	"slices"
)

func (uCase *usecase) CreatePoll(data presenter.PollCreateUpdateRequest) map[string]string {
	var err error

	errMap := make(map[string]string)

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	_, err = uCase.subjectRepo.GetSubjectByID(data.SubjectID)
	if err != nil {
		errMap["subjectID"] = err.Error()
		return errMap
	}

	subjectCourseIDs, err := uCase.subjectRepo.GetCourseIDsBySubjectID(data.SubjectID)

	if err != nil {
		errMap["subjectCourse"] = "error while fetching subject course"

		return errMap
	}

	if !slices.Contains(subjectCourseIDs, data.CourseID) {
		errMap["subjectID"] = fmt.Sprintf("Subject ID %d does not belong to Course ID %d", data.SubjectID, data.CourseID)
		return errMap
	}

	if err = uCase.repo.CreatePoll(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
