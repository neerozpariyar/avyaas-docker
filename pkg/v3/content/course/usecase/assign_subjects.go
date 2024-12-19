package usecase

import "fmt"

func (uCase *usecase) AssignSubjectsToCourse(courseID uint, subjectIDs []uint) map[string]string {
	errMap := make(map[string]string)

	if _, err := uCase.repo.GetCourseByID(courseID); err != nil {

		errMap["course"] = fmt.Sprintf("Course  %d does not Exist", courseID)

		return errMap

	}

	for _, subjectID := range subjectIDs {

		if _, err := uCase.subjectRepo.GetSubjectByID(subjectID); err != nil {

			errMap["subject"] = fmt.Sprintf("Subject  %d does not Exist", subjectID)

			return errMap
		}
	}

	err := uCase.repo.AssignSubjectsToCourse([]uint{courseID}, subjectIDs)

	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
