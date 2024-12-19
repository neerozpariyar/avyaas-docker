package usecase

import (
	"avyaas/internal/domain/presenter"
	"errors"

	"gorm.io/gorm"
)

func (uCase *usecase) GetTestDetails(testID, requesterID uint) (*presenter.TestDetailsPresenter, error) {
	test, err := uCase.repo.GetTestByID(testID)
	if err != nil {
		return nil, err
	}
	user, err := uCase.accountRepo.GetUserByID(requesterID)
	if err != nil {
		return nil, err
	}
	if user.RoleID == 4 {
		if _, err = uCase.repo.GetStudentTest(requesterID, testID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if !*test.IsFree {
					return nil, err
				}
			}
		}
	}

	// testDetails, err := uCase.repo.GetTestDetails(testID, requesterID)
	// if err != nil {
	// 	return testDetails, err
	// }

	testDetails := &presenter.TestDetailsPresenter{
		ID:        testID, // Set the ID field with the actual testID value
		Title:     test.Title,
		StartTime: test.StartTime,
		EndTime:   test.EndTime,
		Duration:  test.Duration,
		ExtraTime: test.ExtraTime,
		Price:     test.Price,
		IsPublic:  test.IsPublic,
		IsPremium: test.IsPremium,
		CreatedBy: test.CreatedBy,
	}

	testTypeData := make(map[string]interface{})

	// Retrieve the test type with the given type ID
	testType, err := uCase.repo.GetTestTypeByID(uint(test.TestTypeID))
	if err != nil {
		return nil, err
	}

	testTypeData["id"] = testType.ID
	testTypeData["title"] = testType.Title

	testDetails.TestType = testTypeData

	// Retrieve the course with the given CourseID
	course, err := uCase.courseRepo.GetCourseByID(test.CourseID)
	if err != nil {
		return nil, err
	}

	courseData := make(map[string]interface{})
	courseData["id"] = course.ID
	courseData["courseID"] = course.CourseID
	testDetails.Course = courseData

	if len(test.QuestionSets) != 0 {
		questionSet, err := uCase.questionSetRepo.GetQuestionSetByID(test.QuestionSets[0].ID)
		if err != nil {
			return nil, err
		}

		testDetails.TotalQuestions = questionSet.TotalQuestions
		testDetails.Marks = questionSet.Marks

		// questionSetPresenter := presenter.QuestionSetDetailsPresenter{
		// 	ID:             questionSet.ID,
		// 	Title:          questionSet.Title,
		// 	Description:    questionSet.Description,
		// 	TotalQuestions: questionSet.TotalQuestions,
		// 	Marks:          questionSet.Marks,
		// 	Course:         courseData,
		// 	File:           questionSet.File,
		// }

		// // var allQuestions []presenter.QuestionPresenter
		// for _, question := range questionSet.Questions {
		// 	// Fetch options for each question
		// 	option, err := uCase.questionRepo.GetOptionByQuestionID(question.ID)
		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	var options []presenter.Option
		// 	if option != nil {
		// 		options = []presenter.Option{}
		// 		if option.OptionA != "" || option.UrlA != "" {
		// 			options = append(options, presenter.Option{
		// 				Title: option.OptionA,
		// 				Url:   utils.GetFileURL(option.UrlA),
		// 			})
		// 		}
		// 		if option.OptionB != "" || option.UrlB != "" {
		// 			options = append(options, presenter.Option{
		// 				Title: option.OptionB,
		// 				Url:   utils.GetFileURL(option.UrlB),
		// 			})
		// 		}
		// 		if option.OptionC != "" || option.UrlC != "" {
		// 			options = append(options, presenter.Option{
		// 				Title: option.OptionC,
		// 				Url:   utils.GetFileURL(option.UrlC),
		// 			})
		// 		}
		// 		if option.OptionD != "" || option.UrlD != "" {
		// 			options = append(options, presenter.Option{
		// 				Title: option.OptionD,
		// 				Url:   utils.GetFileURL(option.UrlD),
		// 			})
		// 		}
		// 	}

		// 	eachQuestion := presenter.QuestionPresenter{
		// 		ID:          question.ID,
		// 		Title:       question.Title,
		// 		Description: question.Description,
		// 		Image:       question.Image,
		// 		ForTest:     question.ForTest,
		// 		Options:     options,
		// 		FileType:    option.FileType,
		// 		Answer:      option.Answer,
		// 	}

		// 	requester, err := uCase.accountRepo.GetUserByID(requesterID)
		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	if requester.RoleID == 4 {
		// 		questionSetBytes, err := json.Marshal(&eachQuestion)
		// 		if err != nil {
		// 			fmt.Println("error marshaling QuestionSet:", err)
		// 		}

		// 		a, err := utils.Encrypt(questionSetBytes)
		// 		if err != nil {
		// 			return nil, err
		// 		}

		// 		encryptedURLString := base64.URLEncoding.EncodeToString(a)

		// 		questionSetPresenter.Questions = append(questionSetPresenter.Questions, encryptedURLString)
		// 	} else {
		// 		questionSetPresenter.Questions = append(questionSetPresenter.Questions, eachQuestion)
		// 	}
		// }

		questionSetDetails, err := uCase.questionSetUsecase.GetQuestionSetDetails(test.QuestionSets[0].ID, requesterID)
		if err != nil {
			return nil, err
		}

		testDetails.QuestionSet = questionSetDetails
	}

	return testDetails, nil
}
