package gorm

// import (
// 	"avyaas/internal/domain/models"
// 	"avyaas/internal/domain/presenter"
// 	"avyaas/utils"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// )

// func (repo *repository) GetTestDetails(testID, requesterID uint) (*presenter.TestDetailsPresenter, error) {
// 	requester, err := repo.accountRepo.GetUserByID(requesterID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var test models.Test
// 	if err := repo.db.Preload("QuestionSets").First(&test, "id = ?", testID).Error; err != nil {
// 		return nil, err
// 	}

// 	testDetails := &presenter.TestDetailsPresenter{
// 		ID:        testID, // Set the ID field with the actual testID value
// 		Title:     test.Title,
// 		StartTime: test.StartTime,
// 		EndTime:   test.EndTime,
// 		Duration:  test.Duration,
// 		ExtraTime: test.ExtraTime,
// 		Price:     test.Price,
// 		IsPublic:  test.IsPublic,
// 		IsPremium: test.IsPremium,
// 		CreatedBy: test.CreatedBy,
// 	}

// 	testData := make(map[string]interface{})

// 	// Retrieve the test type with the given type ID
// 	testType, err := repo.GetTestTypeByID(uint(test.TestTypeID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	testData["id"] = testType.ID
// 	testData["title"] = testType.Title

// 	testDetails.TestType = testData

// 	// Retrieve the course with the given CourseID
// 	course, err := repo.courseRepo.GetCourseByID(test.CourseID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	courseData := make(map[string]interface{})
// 	courseData["id"] = course.ID
// 	courseData["courseID"] = course.CourseID
// 	testDetails.Course = courseData

// 	// Initialize the question set data as nil. If data exists,
// 	testDetails.QuestionSet = nil

// 	var questionSet models.QuestionSet
// 	var questionSetPresenter presenter.QuestionSetDetailsPresenter

// 	if len(test.QuestionSets) != 0 {
// 		if err := repo.db.Preload("Questions").First(&questionSet, "id = ?", test.QuestionSets[0].ID).Error; err != nil {
// 			return nil, err
// 		}

// 		testDetails.TotalQuestions = questionSet.TotalQuestions
// 		testDetails.Marks = questionSet.Marks

// 		questionSetPresenter = presenter.QuestionSetDetailsPresenter{
// 			ID:             questionSet.ID,
// 			Title:          questionSet.Title,
// 			Description:    questionSet.Description,
// 			TotalQuestions: questionSet.TotalQuestions,
// 			Marks:          questionSet.Marks,
// 			Course:         courseData,
// 			File:           questionSet.File,
// 		}

// 		for _, question := range questionSet.Questions {
// 			var qsQuestion *models.QuestionSetQuestion
// 			err := repo.db.Where("question_set_id = ? AND question_id = ?", questionSet.ID, question.ID).First(&qsQuestion).Error
// 			if err != nil {
// 				return nil, err
// 			}

// 			option, err := repo.questionRepo.GetOptionByQuestionID(question.ID)
// 			if err != nil {
// 				return nil, err
// 			}

// 			var options []presenter.Option
// 			if option != nil {
// 				options = []presenter.Option{}
// 				if option.OptionA != "" || option.UrlA != "" {
// 					options = append(options, presenter.Option{
// 						Title: option.OptionA,
// 						Url:   utils.GetFileURL(option.UrlA),
// 					})
// 				}
// 				if option.OptionB != "" || option.UrlB != "" {
// 					options = append(options, presenter.Option{
// 						Title: option.OptionB,
// 						Url:   utils.GetFileURL(option.UrlB),
// 					})
// 				}
// 				if option.OptionC != "" || option.UrlC != "" {
// 					options = append(options, presenter.Option{
// 						Title: option.OptionC,
// 						Url:   utils.GetFileURL(option.UrlC),
// 					})
// 				}
// 				if option.OptionD != "" || option.UrlD != "" {
// 					options = append(options, presenter.Option{
// 						Title: option.OptionD,
// 						Url:   utils.GetFileURL(option.UrlD),
// 					})
// 				}
// 			}

// 			eachQuestion := presenter.QuestionPresenter{
// 				ID:          question.ID,
// 				Title:       question.Title,
// 				Description: question.Description,
// 				Image:       question.Image,
// 				ForTest:     question.ForTest,
// 				Options:     options,
// 				FileType:    option.FileType,
// 				Answer:      option.Answer,
// 			}

// 			subject, err := repo.subjectRepo.GetSubjectByID(question.SubjectID)
// 			if err != nil {
// 				return nil, err
// 			}

// 			course, err := repo.courseRepo.GetCourseByID(subject.CourseID)
// 			if err != nil {
// 				return nil, err
// 			}

// 			courseData := make(map[string]interface{})
// 			courseData["id"] = course.ID
// 			courseData["tile"] = course.Title
// 			eachQuestion.Course = courseData

// 			subjectData := make(map[string]interface{})
// 			subjectData["id"] = subject.ID
// 			subjectData["tile"] = subject.Title
// 			eachQuestion.Subject = subjectData
// 			// for _, option := range options {
// 			// 	eachOption := presenter.Option{
// 			// 		Title: []string{options.OptionA, options.OptionB, options.OptionC, options.OptionD},
// 			// 		Url:   []string{options.UrlA, options.UrlB, options.UrlC, options.UrlD},
// 			// 		FileType:    options.FileType,
// 			// 	Answer:      options.Answer,
// 			// 	}

// 			// 	eachQuestion.Options = append(eachQuestion.Options, eachOption)
// 			// }

// 			if requester.RoleID == 4 {
// 				questionSetBytes, err := json.Marshal(&question)
// 				if err != nil {
// 					fmt.Println("error marshaling QuestionSet:", err)
// 				}

// 				a, err := utils.Encrypt(questionSetBytes)
// 				if err != nil {
// 					return nil, err
// 				}

// 				encryptedURLString := base64.URLEncoding.EncodeToString(a)

// 				questionSetPresenter.Questions = append(questionSetPresenter.Questions, encryptedURLString)
// 			} else {
// 				questionSetPresenter.Questions = append(questionSetPresenter.Questions, question)
// 			}
// 		}

// 		testDetails.QuestionSet = &questionSetPresenter
// 	}
// 	return testDetails, nil
// }
