package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

/*
CreateQuestionSet is a usecase method responsible for creating a new question set.

Parameters:
  - data: A models.QuestionSet instance representing the question set to be created.

Returns:
  - errMap: A map[string]string containing error messages, if any, encountered during the process.
*/
func (uCase *usecase) CreateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	course, err := uCase.courseRepo.GetCourseByID(data.CourseID)
	if err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	if questionSet, err := uCase.repo.GetQuestionSetByTitleAndCourseID(data.Title, data.CourseID); err == nil {
		if questionSet.CourseID == data.CourseID {
			errMap["title"] = fmt.Errorf("question set with title '%s' already exists in course '%s'", data.Title, course.Title).Error()
			return errMap
		}
	}

	if err = uCase.repo.CreateQuestionSet(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
