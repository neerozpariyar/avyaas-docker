package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

/*
CreateQuestion is a usecase method responsible for creating a new question.

Parameters:
  - data: A models.Question instance representing the question to be created.

Returns:
  - errMap: A map[string]string containing error messages, if any, encountered during the process.
*/

func (uCase *usecase) CreateTypeQuestion(data presenter.TypeQuestionPresenter) map[string]string {
	errMap := make(map[string]string)
	// Validate Subject and QuestionSet IDs
	subject, err := uCase.subjectRepo.GetSubjectByID(data.SubjectID)
	if err != nil {
		errMap["subjectID"] = err.Error()
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

	switch data.Type {
	case "CaseBased":
		_, err := uCase.repo.CreateCaseTypeQuestion(data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	case "FillInTheBlanks":
		err := uCase.repo.CreateFillInBlanksQuestion(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	case "MCQ", "MultiAnswer":
		err := uCase.repo.CreateMCQQuestion(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	case "TrueFalse":
		err := uCase.repo.CreateTrueOrFalseQuestion(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	default:
		errMap["error"] = "Invalid question type"
		return errMap
	}

	// if _, err := uCase.repo.CreateTypeQuestion(data); err != nil {
	// 	errMap["error"] = err.Error()
	// 	return errMap
	// }

	return errMap
}
