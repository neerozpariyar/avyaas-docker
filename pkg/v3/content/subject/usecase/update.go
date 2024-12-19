package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) UpdateSubject(data presenter.SubjectCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing subject  with the provided subject's ID
	sub, err := uCase.repo.GetSubjectByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}
	for _, courseID := range data.CourseIDs {
		if _, err := uCase.courseRepo.GetCourseByID(courseID); err != nil {
			errMap["courseID"] = err.Error()
			return errMap
		}
	}
	// Check if a subject  with the given SubjectID already exists
	subByID, err := uCase.repo.GetSubjectBySubjectID(data.SubjectID)
	if err == nil {
		// Check if the subjectID is the same as of the requested subject
		if sub.SubjectID != subByID.SubjectID {
			errMap["subjectID"] = fmt.Errorf("subject  with  id: '%s' already exists", subByID.SubjectID).Error()
			return errMap
		}
	}

	// Delegate the update of subject
	if err = uCase.repo.UpdateSubject(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
