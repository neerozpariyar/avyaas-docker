package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func (uCase *usecase) GetQuestionSetDetails(id, requesterID uint) (*presenter.QuestionSetDetailsPresenter, error) {
	questionSet, err := uCase.repo.GetQuestionSetByID(id)
	if err != nil {
		return nil, err
	}

	course, err := uCase.courseRepo.GetCourseByID(questionSet.CourseID)
	if err != nil {
		return nil, err
	}
	var allCourses []presenter.CourseDataForPoll
	courseData := make(map[string]interface{})
	courseData["id"] = course.ID
	courseData["courseID"] = course.CourseID

	questionSetPresenter := presenter.QuestionSetDetailsPresenter{
		ID:             questionSet.ID,
		Title:          questionSet.Title,
		Description:    questionSet.Description,
		TotalQuestions: questionSet.TotalQuestions,
		Marks:          questionSet.Marks,
		Course:         courseData,
		File:           questionSet.File,
	}

	for i, question := range questionSet.Questions {
		// Fetch options for each question
		// option, err := uCase.questionRepo.GetTypeOptionByQuestionID(question.ID)
		// if err != nil {
		// 	return nil, err
		// }

		options, err := uCase.questionRepo.GetOptionsByQuestionID(question.ID)
		if err != nil {
			return nil, err
		}

		// var options []presenter.TypeOptionListPresenter

		eachQuestion := presenter.QuestionListResponse{
			ID:      question.ID,
			Title:   question.Title,
			ForTest: question.ForTest,
			Options: make([]presenter.OptionListPresenter, len(options)),
			IsTrue:  question.IsTrue,
		}
		bookmark, isBookmarked, err := uCase.bookmarkRepo.GetBookmarkedQuestionAndCheckIfBookmarked(requesterID, question.ID)
		if err != nil {
			return nil, err
		}
		eachQuestion.IsBookmarked = isBookmarked

		if isBookmarked {
			eachQuestion.BookmarkID = bookmark.ID
		}
		if questionSet.Questions[i].Type == "CaseBased" {
			nestedQuestions, err := uCase.questionRepo.GetNestedQuestions(question.ID)
			if err != nil {
				return nil, err
			}

			for i := range nestedQuestions {
				nestedQuestions[i].IsBookmarked = eachQuestion.IsBookmarked
				nestedQuestions[i].BookmarkID = eachQuestion.BookmarkID
			}
			eachQuestion.CaseQuestionID = question.ID
			eachQuestion.Questions = nestedQuestions
		}
		for i, opt := range options {

			var audio, image string

			if opt.Audio != "" {
				audio = utils.GetFileURL(opt.Audio)
			}

			if opt.Image != "" {
				image = utils.GetFileURL(opt.Image)
			}
			eachQuestion.Options[i] = presenter.OptionListPresenter{
				ID:        opt.ID,
				Text:      opt.Text,
				Audio:     &audio,
				Image:     &image,
				IsCorrect: &opt.IsCorrect,
			}
		}

		// if question.Image != "" {
		// 	eachQuestion.Image = utils.GetFileURL(question.Image)
		// }

		// if question.Audio != "" {
		// 	eachQuestion.Audio = utils.GetFileURL(question.Audio)
		// }

		subject, err := uCase.subjectRepo.GetSubjectByID(question.SubjectID)
		if err != nil {
			return nil, err
		}

		courses, err := uCase.subjectRepo.GetCoursesBySubjectId(id)

		if err != nil {
			return nil, err
		}

		for _, course := range courses {
			singleCourseData := presenter.CourseDataForPoll{
				ID:       course.ID,
				Title:    course.Title,
				CourseID: course.CourseID,
			}

			allCourses = append(allCourses, singleCourseData)
		}

		eachQuestion.Courses = allCourses

		subjectData := make(map[string]interface{})
		subjectData["id"] = subject.ID
		subjectData["title"] = subject.Title
		eachQuestion.Subject = subjectData

		requester, err := uCase.accountRepo.GetUserByID(requesterID)
		if err != nil {
			return nil, err
		}

		if requester.RoleID == 4 {
			questionSetBytes, err := json.Marshal(&eachQuestion)
			if err != nil {
				fmt.Println("error marshaling QuestionSet:", err)
			}

			a, err := utils.Encrypt(questionSetBytes)
			if err != nil {
				return nil, err
			}

			encryptedURLString := base64.URLEncoding.EncodeToString(a)

			questionSetPresenter.Questions = append(questionSetPresenter.Questions, encryptedURLString)
		} else {
			questionSetPresenter.Questions = append(questionSetPresenter.Questions, eachQuestion)
		}

	}

	return &questionSetPresenter, nil
}

// func (uCase *usecase) GetQuestionSetDetails(id uint) (*presenter.QuestionSetDetailsPresenter, error) {
// 	questionSet, err := uCase.repo.GetQuestionSetByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	course, err := uCase.courseRepo.GetCourseByID(questionSet.CourseID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	courseData := make(map[string]interface{})
// 	courseData["id"] = course.ID
// 	courseData["courseID"] = course.CourseID

// 	questionSetPresenter := presenter.QuestionSetDetailsPresenter{
// 		ID:             questionSet.ID,
// 		Title:          questionSet.Title,
// 		Description:    questionSet.Description,
// 		TotalQuestions: questionSet.TotalQuestions,
// 		Marks:          questionSet.Marks,
// 		Course:         courseData,
// 		File:           questionSet.File,
// 	}

// 	for _, question := range questionSet.Questions {
// 		qsQuestion, err := uCase.repo.GetQuestionSetQuestion(id, question.ID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		subject, err := uCase.subjectRepo.GetSubjectByID(question.SubjectID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		subjectData := make(map[string]interface{})
// 		subjectData["id"] = subject.ID
// 		subjectData["title"] = subject.Title

// 		questionPresenter := presenter.QuestionDetailsPresenter{
// 			ID:          question.ID,
// 			Title:       question.Title,
// 			Description: question.Description,
// 			Image:       question.Image,
// 			Position:    qsQuestion.Position,
// 			ForTest:     question.ForTest,
// 			Subject:     subjectData,
// 		}

// 		option, err := uCase.questionRepo.GetOptionByQuestionID(question.ID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		optionPresenter := presenter.OptionDetailsPresenter{
// 			ID:       option.ID,
// 			OptionA:  option.OptionA,
// 			OptionB:  option.OptionB,
// 			OptionC:  option.OptionC,
// 			OptionD:  option.OptionD,
// 			UrlA:     option.UrlA,
// 			UrlB:     option.UrlB,
// 			UrlC:     option.UrlC,
// 			UrlD:     option.UrlD,
// 			Answer:   option.Answer,
// 			FileType: option.FileType,
// 		}

// 		questionPresenter.Options = optionPresenter

// 		questionSetPresenter.Questions = append(questionSetPresenter.Questions, questionPresenter)

// 	}

// 	return &questionSetPresenter, nil
// }
