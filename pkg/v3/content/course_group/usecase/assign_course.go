package usecase

import "fmt"

func (uCase *usecase) AssignCoursesToCourseGroup(courseGroupID uint, courseIds []uint) map[string]string {
	errMap := make(map[string]string)

	if _, err := uCase.repo.GetCourseGroupByID(courseGroupID); err != nil {
		errMap["courseGroup"] = fmt.Sprintf("Course Group  %d does not Exist", courseGroupID)
		return errMap
	}

	for _, courseId := range courseIds {

		if _, err := uCase.courseRepo.GetCourseByID(courseId); err != nil {
			errMap["course"] = fmt.Sprintf("Course  %d does not Exist", courseId)
			return errMap
		}

	}

	err := uCase.repo.AssignCoursesToCourseGroup([]uint{courseGroupID}, courseIds)

	if err != nil {
		errMap["assigning"] = err.Error()
		return errMap
	}

	return nil
}
