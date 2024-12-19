package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
	"slices"
)

func (uCase *usecase) UpdateDiscussion(discussion presenter.DiscussionCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing discussion  with the provided discussion 's ID
	_, err = uCase.repo.GetDiscussionByID(discussion.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}
	if _, err := uCase.courseRepo.GetCourseByID(discussion.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	_, err = uCase.subjectRepo.GetSubjectByID(discussion.SubjectID)
	if err != nil {
		errMap["subjectID"] = err.Error()
		return errMap
	}

	if subjectCourseIDs, err := uCase.subjectRepo.GetCourseIDsBySubjectID(discussion.SubjectID); err != nil || !slices.Contains(subjectCourseIDs, discussion.CourseID) {
		errMap["subjectID"] = fmt.Sprintf("Subject ID %d does not belong to Course ID %d", discussion.SubjectID, discussion.CourseID)
		return errMap
	}
	// Delegate the update of discussion
	if err = uCase.repo.UpdateDiscussion(discussion); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
