package usecase

import (
	"avyaas/internal/domain/presenter"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (uCase *usecase) CreateCourse(data presenter.CourseCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	for _, CourseGroupID := range data.CourseGroupIDs {
		if _, err := uCase.courseGroupRepo.GetCourseGroupByID(CourseGroupID); err != nil {
			errMap["courseGroupID"] = err.Error()
			return errMap
		}
	}

	// Check if a course with the provided courseID already exists
	_, err = uCase.repo.GetCourseByCourseID(data.CourseID)
	if err == nil {
		// If a course with the given courseID already exists, return an error
		errMap["courseID"] = fmt.Errorf("course with id: '%s' already exists", data.CourseID).Error()
		return errMap
	}

	//If no course is found with the provided courseID, create the course
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = uCase.repo.CreateCourse(data); err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	}

	return errMap

}

// // Check if a course with the provided courseID already exists
// _, err = uCase.repo.GetCourseByCourseGroupID(presenter.courseID)
// if err == nil {
// 	// If a course with the given courseID already exists, return an error
// 	errMap["groupID"] = fmt.Errorf("course with id: '%s' already exists", presenter.courseID).Error()
// 	return errMap
// }

// // If no course is found with the provided courseIsD, create the course group
// if errors.Is(err, gorm.ErrRecordNotFound) {
// 	if err = uCase.repo.CreateCourse(course); err != nil {
// 		errMap["error"] = err.Error()
// 		return errMap
// 	}
// }

// return errMap
