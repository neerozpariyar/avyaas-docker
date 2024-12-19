package usecase

import (
	"avyaas/internal/domain/presenter"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
CreateCourseGroup is a usecase method responsible for creating a new course group.

Parameters:
  - courseGroup: A models.CourseGroup instance representing the course group to be created.

Returns:
  - errMap: A map[string]string containing error messages, if any, encountered during the process.
*/
func (uCase *usecase) CreateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Check if a course group with the provided GroupID already exists
	_, err = uCase.repo.GetCourseGroupByGroupID(data.GroupID)
	if err == nil {
		// If a course group with the given GroupID already exists, return an error
		errMap["groupID"] = fmt.Errorf("course group with group id: '%s' already exists", data.GroupID).Error()
		return errMap
	}

	// If no course group is found with the provided GroupID, create the course group
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = uCase.repo.CreateCourseGroup(data); err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	}

	return errMap

}
