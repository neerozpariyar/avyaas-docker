package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
)

func (uCase *usecase) GetTestResult(testID, requesterID uint) (*presenter.TestResultResponse, error) {
	var result *presenter.TestResultResponse

	test, err := uCase.repo.GetTestByID(testID)
	if err != nil {
		return nil, err
	}

	studentTest, err := uCase.repo.GetStudentTest(requesterID, testID)
	if err != nil {
		return nil, err
	}

	if !studentTest.HasAttended {
		return nil, fmt.Errorf("test not attended yet")
	}

	responses, err := uCase.repo.GetTestResult(testID, requesterID)
	if err != nil {
		return nil, err
	}

	if len(responses) == 0 {
		return result, nil
	}

	testResult, err := uCase.repo.GetStudentTestResult(requesterID, testID)
	if err != nil {
		return nil, err
	}

	result = &presenter.TestResultResponse{
		ID:               test.ID,
		Title:            test.Title,
		TotalQuestions:   test.QuestionSets[0].TotalQuestions,
		Marks:            test.QuestionSets[0].Marks,
		Score:            int(testResult.Score),
		TotalAttempted:   testResult.TotalAttempted,
		TotalUnattempted: testResult.TotalUnattempted,
		TotalCorrect:     testResult.TotalCorrect,
		TotalWrong:       testResult.TotalWrong,
	}

	testType, err := uCase.repo.GetTestTypeByID(uint(test.TestTypeID))
	if err != nil {
		return nil, err
	}

	result.Type = testType.Title

	course, err := uCase.courseRepo.GetCourseByID(test.CourseID)
	if err != nil {
		return nil, err
	}

	result.Course = course.Title

	var questions []presenter.TestResultQuestionPresenter
	for _, response := range responses {
		question, err := uCase.questionRepo.GetQuestionByID(response.QuestionID)
		if err != nil {
			return nil, err
		}

		questionPresenter := presenter.TestResultQuestionPresenter{
			ID:    question.ID,
			Title: question.Title,
			Image: question.Image,
		}

		options, err := uCase.questionRepo.GetOptionsByQuestionID(question.ID)
		if err != nil {
			return nil, err
		}
		var option presenter.OptionListPresenter
		for _, opt := range options {

			var audio, image string

			if opt.Image != "" {
				audio = utils.GetFileURL(opt.Audio)
			}

			if opt.Image != "" {
				image = utils.GetFileURL(opt.Image)
			}

			option = presenter.OptionListPresenter{
				ID:        opt.ID,
				Text:      opt.Text,
				Audio:     &audio,
				Image:     &image,
				IsCorrect: &opt.IsCorrect,
			}
		}
		questionPresenter.Options = option

		questionPresenter.SelectedOptionID = response.AnswerID

		questions = append(questions, questionPresenter)
	}

	result.Questions = questions

	return result, nil
}
