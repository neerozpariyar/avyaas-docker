package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
	"slices"
)

func (uCase *usecase) CreateDiscussion(data presenter.DiscussionCreateUpdateRequest) map[string]string {

	errMap := make(map[string]string)

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	_, err := uCase.subjectRepo.GetSubjectByID(data.SubjectID)
	if err != nil {
		errMap["subjectID"] = err.Error()
		return errMap
	}
	if subjectCourseIDs, err := uCase.subjectRepo.GetCourseIDsBySubjectID(data.SubjectID); err != nil || !slices.Contains(subjectCourseIDs, data.CourseID) {
		errMap["subjectID"] = fmt.Sprintf("Subject ID %d does not belong to Course ID %d", data.SubjectID, data.CourseID)
		return errMap
	}
	if err = uCase.repo.CreateDiscussion(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
