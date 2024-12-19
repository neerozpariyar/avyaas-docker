package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

/*
ListQuestion retrieves a paginated list of questions from the repository.

Parameters:
  - page: An integer representing the page number for pagination.

Returns:
  - questions: A slice of models.Question representing the retrieved questions.
  - totalPage: An integer representing the total number of pages available.
  - error: An error indicating the success or failure of the operation.
*/
func (u *usecase) ListQuestion(request presenter.ListQuestionRequest) ([]presenter.QuestionListResponse, int, error) {
	// Delegate the retrieval of questions
	questions, totalPage, err := u.repo.ListQuestion(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allQuestions []presenter.QuestionListResponse
	var allCourses []presenter.CourseDataForPoll

	for _, question := range questions {
		option, err := u.repo.GetOptionsByQuestionID(question.ID)
		if err != nil {
			return nil, 0, err
		}

		var options []presenter.OptionListPresenter
		for i, opt := range option {
			var image, audio string
			if option[i].Image != "" {
				image = utils.GetFileURL(option[i].Image)
			}
			if option[i].Audio != "" {
				audio = utils.GetFileURL(option[i].Audio)
			}
			options = append(options, presenter.OptionListPresenter{
				ID:         opt.ID,
				QuestionID: opt.QuestionID,
				Image:      &image,
				Audio:      &audio,
				Text:       opt.Text,
				IsCorrect:  &opt.IsCorrect})

			if question.Type == "FillInTheBlanks" {
				options[i].Audio = nil
				options[i].Image = nil
				options[i].IsCorrect = nil
			}
		}

		eachQuestion := presenter.QuestionListResponse{
			ID:           question.ID,
			Title:        question.Title,
			Description:  *question.Description,
			Type:         question.Type,
			ForTest:      question.ForTest,
			NegativeMark: question.NegativeMark,
			IsTrue:       question.IsTrue,
			Options:      options,
			// QuestionSetID: *question.QuestionSetID,
		}
		bookmark, isBookmarked, err := u.bookmarkRepo.GetBookmarkedQuestionAndCheckIfBookmarked(request.RequesterID, question.ID)
		if err != nil {
			return nil, 0, err
		}
		eachQuestion.IsBookmarked = isBookmarked

		if isBookmarked {
			eachQuestion.BookmarkID = bookmark.ID
		}
		switch question.Type {
		case "CaseBased":
			nestedQuestions, err := u.repo.GetNestedQuestions(question.ID)
			if err != nil {
				return nil, 0, err
			}
			for i := range nestedQuestions {
				nestedQuestions[i].IsBookmarked = eachQuestion.IsBookmarked
				nestedQuestions[i].BookmarkID = eachQuestion.BookmarkID
			}
			eachQuestion.CaseQuestionID = question.ID
			eachQuestion.Questions = nestedQuestions
		case "MCQ", "MultiAnswer":
			eachQuestion.IsTrue = nil
		case "TrueFalse":
			eachQuestion.Description = ""

			eachQuestion.IsTrue = question.IsTrue
		case "FillInTheBlanks":
			eachQuestion.Description = ""

		}
		subject, err := u.subjectRepo.GetSubjectByID(question.SubjectID)
		if err != nil {
			return nil, 0, err
		}

		courses, err := u.subjectRepo.GetCoursesBySubjectId(subject.ID)
		if err != nil {
			return nil, 0, err
		}

		for _, course := range courses {
			singleCourse := presenter.CourseDataForPoll{
				ID:       course.ID,
				Title:    course.Title,
				CourseID: course.CourseID,
			}

			allCourses = append(allCourses, singleCourse)
		}

		eachQuestion.Courses = allCourses

		subjectData := make(map[string]interface{})
		subjectData["id"] = subject.ID
		subjectData["title"] = subject.Title
		eachQuestion.Subject = subjectData

		allQuestions = append(allQuestions, eachQuestion)
	}

	return allQuestions, int(totalPage), nil
}

// for _, question := range questions {
// 	// Fetch options for each question
// 	option, err := u.repo.GetOptionsByQuestionID(question.ID)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	var options []presenter.Option
// 	if option != nil {
// 		optionA := presenter.Option{
// 			Title: option.OptionA,
// 		}

// 		if option.ImageA != "" {
// 			optionA.Image = utils.GetFileURL(option.ImageA)
// 		}

// 		if option.AudioA != "" {
// 			optionA.Audio = utils.GetFileURL(option.AudioA)
// 		}
// 		options = append(options, optionA)

// 		optionB := presenter.Option{
// 			Title: option.OptionB,
// 		}

// 		if option.ImageB != "" {
// 			optionB.Image = utils.GetFileURL(option.ImageB)
// 		}

// 		if option.AudioB != "" {
// 			optionB.Audio = utils.GetFileURL(option.AudioB)
// 		}
// 		options = append(options, optionB)

// 		optionC := presenter.Option{
// 			Title: option.OptionC,
// 		}

// 		if option.ImageC != "" {
// 			optionC.Image = utils.GetFileURL(option.ImageC)
// 		}

// 		if option.AudioC != "" {
// 			optionC.Audio = utils.GetFileURL(option.AudioC)
// 		}
// 		options = append(options, optionC)

// 		optionD := presenter.Option{
// 			Title: option.OptionD,
// 		}

// 		if option.ImageD != "" {
// 			optionD.Image = utils.GetFileURL(option.ImageD)
// 		}

// 		if option.AudioD != "" {
// 			optionD.Audio = utils.GetFileURL(option.AudioD)
// 		}
// 		options = append(options, optionD)
// 	}

// 	eachQuestion := presenter.QuestionPresenter{
// 		ID:      question.ID,
// 		Title:   question.Title,
// 		ForTest: question.ForTest,
// 		Options: options,
// 		Answer:  option.Answer,
// 	}

// 	if question.Image != "" {
// 		eachQuestion.Image = utils.GetFileURL(question.Image)
// 	}

// 	if question.Audio != "" {
// 		eachQuestion.Audio = utils.GetFileURL(question.Audio)
// 	}

// 	bookmark, isBookmarked, err := u.bookmarkRepo.GetBookmarkedQuestionAndCheckIfBookmarked(request.RequesterID, question.ID)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	eachQuestion.IsBookmarked = isBookmarked

// 	if isBookmarked {
// 		eachQuestion.BookmarkID = bookmark.ID
// 	}

// 	subject, err := u.subjectRepo.GetSubjectByID(question.SubjectID)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	course, err := u.courseRepo.GetCourseByID(subject.CourseID)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	courseData := make(map[string]interface{})
// 	courseData["id"] = course.ID
// 	courseData["title"] = course.Title
// 	eachQuestion.Course = courseData

// 	subjectData := make(map[string]interface{})
// 	subjectData["id"] = subject.ID
// 	subjectData["title"] = subject.Title
// 	eachQuestion.Subject = subjectData

// 	allQuestions = append(allQuestions, eachQuestion)
// }
