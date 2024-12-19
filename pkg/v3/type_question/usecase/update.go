package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

/*
UpdateQuestion is a usecase method responsible for updating a new question.

Parameters:
  - question: A models.Question instance representing the question to be updated.

Returns:
  - errMap: A map[string]string containing error messages, if any, encountered during the process.
*/
func (uCase *usecase) UpdateTypeQuestion(data presenter.TypeQuestionPresenter) map[string]string {
	var err error
	errMap := make(map[string]string)

	_, err = uCase.repo.GetTypeQuestionByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	_, err = uCase.repo.GetTypeOptionsByQuestionID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	subject, err := uCase.subjectRepo.GetSubjectByID(data.SubjectID)
	if err != nil {
		errMap["subjectID"] = err.Error()
		return errMap
	}

	ques, err := uCase.repo.GetTypeQuestionByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if data.QuestionSetID != 0 {
		questionSet, err := uCase.questionSetRepo.GetQuestionSetByID(data.QuestionSetID)
		if err != nil {
			errMap["questionSetID"] = err.Error()
			return errMap
		}

		// check if the question is being assigned to the subject with courseID as same as of the question set
		if questionSet.CourseID != subject.CourseID {
			errMap["error"] = fmt.Errorf("unable to assign question '%s' to question set '%s' of different course", data.Title, questionSet.Title).Error()
			return errMap
		}
	}
	switch ques.Type {
	case "CaseBased":
		_, err := uCase.repo.UpdateCaseTypeQuestion(data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	case "FillInTheBlanks":
		err := uCase.repo.UpdateFillInBlanksQuestion(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	case "MCQ", "MultiAnswer":

		err := uCase.repo.UpdateMCQQuestion(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	case "TrueFalse":
		err := uCase.repo.UpdateTrueOrFalseQuestion(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	default:
		errMap["error"] = "Invalid question type"
		return errMap
	}

	return errMap
}
