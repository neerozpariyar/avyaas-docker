package usecase

import (
	"avyaas/internal/domain/presenter"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (uCase *usecase) CreateSubject(data presenter.SubjectCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	for _, courseID := range data.CourseIDs {
		if _, err := uCase.courseRepo.GetCourseByID(courseID); err != nil {
			errMap["courseID"] = err.Error()
			return errMap
		}
	}

	// Check if a subject group with the provided subjectID already exists
	_, err = uCase.repo.GetSubjectBySubjectID(data.SubjectID)
	if err == nil {
		// If a subject group with the given subjectID already exists, return an error
		errMap["subjectID"] = fmt.Errorf("subject with  subjectID: '%s' already exists", data.SubjectID).Error()
		return errMap
	}
	//If no subject group is found with the provided subjectID, create the subject group
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = uCase.repo.CreateSubject(data); err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	}

	return errMap

}
