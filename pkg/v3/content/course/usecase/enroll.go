package usecase

func (uCase *usecase) EnrollInCourse(userID, courseID uint) map[string]string {
	var err error
	errMap := make(map[string]string)

	course, err := uCase.repo.GetCourseByID(courseID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if !*course.Available {
		errMap["error"] = "course is not available for enrollment"
		return errMap
	}

	_, err = uCase.repo.CheckStudentCourse(userID, courseID)
	if err == nil {
		errMap["error"] = "already enrolled in the course"
		return errMap
	}

	if err = uCase.repo.EnrollInCourse(userID, courseID); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
