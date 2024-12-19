package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (usecase *usecase) ListTypeQuestion(request presenter.ListQuestionRequest) ([]presenter.TypeQuestionListPresenter, int, error) {
	typeQuestions, totalPage, err := usecase.repo.ListTypeQuestion(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allQuestions []presenter.TypeQuestionListPresenter

	for _, question := range typeQuestions {
		option, err := usecase.repo.GetTypeOptionsByQuestionID(question.ID)
		if err != nil {
			return nil, 0, err
		}

		var options []presenter.TypeOptionListPresenter
		for i, opt := range option {
			var image, audio string
			if option[i].Image != nil && *option[i].Image != "" {
				image = utils.GetFileURL(*option[i].Image)
			}
			if option[i].Audio != nil && *option[i].Audio != "" {
				audio = utils.GetFileURL(*option[i].Audio)
			}
			options = append(options, presenter.TypeOptionListPresenter{
				ID:         opt.ID,
				QuestionID: opt.QuestionID,
				Image:      &image,
				Audio:      &audio,
				Text:       *opt.Text,
				IsCorrect:  &opt.IsCorrect})

			if question.Type == "FillInTheBlanks" {
				options[i].Audio = nil
				options[i].Image = nil
				options[i].IsCorrect = nil
			}
		}

		eachQuestion := presenter.TypeQuestionListPresenter{
			ID:           question.ID,
			Title:        question.Title,
			Description:  question.Description,
			Type:         question.Type,
			ForTest:      question.ForTest,
			SubjectID:    question.SubjectID,
			NegativeMark: question.NegativeMark,
			IsTrue:       question.IsTrue,
			Options:      options,
			// QuestionSetID: *question.QuestionSetID,
		}
		bookmark, isBookmarked, err := usecase.bookmarkRepo.GetBookmarkedQuestionAndCheckIfBookmarked(request.RequesterID, question.ID)
		if err != nil {
			return nil, 0, err
		}
		eachQuestion.IsBookmarked = isBookmarked

		if isBookmarked {
			eachQuestion.BookmarkID = bookmark.ID
		}
		switch question.Type {
		case "CaseBased":
			nestedQuestions, err := usecase.repo.GetNestedQuestions(question.ID)
			if err != nil {
				return nil, 0, err
			}
			for i := range nestedQuestions {
				nestedQuestions[i].IsBookmarked = eachQuestion.IsBookmarked
				nestedQuestions[i].BookmarkID = eachQuestion.BookmarkID
			}
			eachQuestion.CaseQuestionID = &question.ID
			eachQuestion.Questions = nestedQuestions
		case "MCQ", "MultiAnswer":
			eachQuestion.IsTrue = nil
			eachQuestion.Description = nil
		case "TrueFalse":
			eachQuestion.Description = nil

			eachQuestion.IsTrue = question.IsTrue
		case "FillInTheBlanks":
			eachQuestion.Description = nil

		}
		subject, err := usecase.subjectRepo.GetSubjectByID(question.SubjectID)
		if err != nil {
			return nil, 0, err
		}

		course, err := usecase.courseRepo.GetCourseByID(subject.CourseID)
		if err != nil {
			return nil, 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["title"] = course.Title
		eachQuestion.Course = courseData

		subjectData := make(map[string]interface{})
		subjectData["id"] = subject.ID
		subjectData["title"] = subject.Title
		eachQuestion.Subject = subjectData

		allQuestions = append(allQuestions, eachQuestion)
	}

	return allQuestions, int(totalPage), nil
}
