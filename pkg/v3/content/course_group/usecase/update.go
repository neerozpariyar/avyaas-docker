package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

/*
UpdateCourseGroup is a usecase method for updating the course group in the repository.

Parameters:
  - courseGroup: A models.CourseGroup struct containing the updated details of the course group.

Returns:
  - errMap: A map containing error messages, if any, encountered during the update operation.
*/
func (uCase *usecase) UpdateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing course group with the provided course group's ID
	cg, err := uCase.repo.GetCourseGroupByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Check if a course group with the given GroupID already exists
	cgByID, err := uCase.repo.GetCourseGroupByGroupID(data.GroupID)
	if err == nil {
		// Check if the groupID is the same as of the requested course group
		if cg.GroupID != cgByID.GroupID {
			errMap["groupID"] = fmt.Errorf("course group with group id: '%s' already exists", cgByID.GroupID).Error()
			return errMap
		}
	}

	// Delegate the update of course group
	if err = uCase.repo.UpdateCourseGroup(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
