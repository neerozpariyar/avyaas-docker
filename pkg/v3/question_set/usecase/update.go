package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

/*
UpdateQuestionSet is a usecase method for updating the question set in the repository.

Parameters:
  - questionSet: A models.QuestionSet struct containing the updated details of the question set.

Returns:
  - errMap: A map containing error messages, if any, encountered during the update operation.
*/
func (uCase *usecase) UpdateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing question set with the provided question set's ID
	_, err = uCase.repo.GetQuestionSetByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	course, err := uCase.courseRepo.GetCourseByID(data.CourseID)
	if err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	if questionSet, err := uCase.repo.GetQuestionSetByTitleAndCourseID(data.Title, data.CourseID); err == nil {
		if questionSet.ID != data.ID {
			errMap["title"] = fmt.Errorf("question set with title '%s' already exists in course '%s'", data.Title, course.Title).Error()
			return errMap
		}
	}

	// Delegate the update of question set
	if err = uCase.repo.UpdateQuestionSet(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
