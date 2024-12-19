package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) GetBookmarkDetails(id uint) (*presenter.BookmarkDetailResponse, map[string]string) {
	errMap := make(map[string]string)
	response := &presenter.BookmarkDetailResponse{}

	bookmark, err := u.repo.GetBookmarkByID(id)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	bookmarkType, err := u.repo.GetBookmarkTypeByID(id)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	if bookmarkType == "content" {
		content, err := u.contentRepo.GetContentByID(bookmark.ContentID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		response.ID = bookmark.ID
		response.ContentID = bookmark.ContentID
		response.Content = content
		response.Question = nil // Assign nil instead of &models.Question{}
	} else if bookmarkType == "question" {
		question, err := u.questionRepo.GetQuestionByID(bookmark.QuestionID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		if question.Image != "" {
			question.Image = utils.GetFileURL(question.Image)
		}
		if question.Audio != "" {
			question.Audio = utils.GetFileURL(question.Audio)
		}
		// Fetch options for each question
		// option, err := uCase.questionRepo.GetTypeOptionByQuestionID(question.ID)
		// if err != nil {
		// 	return nil, err
		// }

		options, err := u.questionRepo.GetOptionsByQuestionID(question.ID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		// var options []presenter.TypeOptionListPresenter

		eachQuestion := presenter.QuestionListResponse{
			ID:    question.ID,
			Title: question.Title,

			Options: make([]presenter.OptionListPresenter, len(options)),
			IsTrue:  question.IsTrue,
		}
		if question.Type == "CaseBased" {
			nestedQuestions, err := u.questionRepo.GetNestedQuestions(question.ID)
			if err != nil {
				errMap["error"] = err.Error()
				return nil, errMap
			}

			eachQuestion.CaseQuestionID = question.ID
			eachQuestion.Questions = nestedQuestions
		}
		var audio, image string

		for i, opt := range options {

			if opt.Image != "" {
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

		response.ID = bookmark.ID
		response.QuestionID = bookmark.QuestionID
		questionData := make(map[string]interface{})
		if question.Type == "CaseBased" {
			questionData["question"] = question
			questionData["nestedQuestion"] = eachQuestion.Questions
		} else {

			questionData["question"] = question
			questionData["options"] = eachQuestion.Options
		}

		response.Question = questionData
		response.Content = nil
	}

	return response, nil
}
