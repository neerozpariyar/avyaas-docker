package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"

	// "avyaas/pkg/v3/account/repository/gorm"
	"avyaas/utils"
	"avyaas/utils/file"
	"errors"
)

func (r *Repository) CreateMCQQuestion(question *presenter.CreateUpdateQuestionRequest) error {
	tsx := r.db.Begin()
	var image, audio string
	isActive := true

	if question.Image != nil {
		imageData, err := file.UploadFile("question", question.Image)
		if err != nil {
			tsx.Rollback()

			return err
		}
		image = utils.GetURLObject(imageData.Url)

		err = tsx.Create(&models.File{
			Title:    imageData.Filename,
			Type:     imageData.FileType,
			Url:      image,
			IsActive: &isActive,
		}).Error

		if err != nil {
			tsx.Rollback()
			return err
		}
	}

	if question.Audio != nil {
		audioData, err := file.UploadFile("question", question.Audio)
		if err != nil {
			tsx.Rollback()
			return err
		}
		audio = utils.GetURLObject(audioData.Url)
		err = tsx.Create(&models.File{
			Title:    audioData.Filename,
			Type:     audioData.FileType,
			Url:      audio,
			IsActive: &isActive,
		}).Error

		if err != nil {
			tsx.Rollback()
			return err
		}
	}
	questionModel := models.Question{
		Timestamp:    models.Timestamp{ID: question.ID},
		Title:        question.Title,
		Image:        image,
		Audio:        audio,
		Type:         question.Type,
		ForTest:      question.ForTest,
		SubjectID:    question.SubjectID,
		NegativeMark: question.NegativeMark,
		Description:  question.Description,
		// QuestionSetID: &question.QuestionSetID,
	}
	if question.NestedQuestionType == "CaseBased" {
		questionModel.CaseQuestionID = question.CaseQuestionID
	}

	// Check if only one option is marked as correct for MCQ type questions
	// and multiple options can be marked as correct for MultiAnswer type questions
	correctCount := 0

	for _, option := range question.Options {
		if option.IsCorrect {
			correctCount++
		}
	}

	if question.Type == "MCQ" {

		if correctCount != 1 {
			return errors.New("MCQ  should have exactly one correct answer")
		}
	} else {
		if correctCount < 2 {
			return errors.New("MultiAnswer  should have more than one correct answer")
		}
	}
	if err := tsx.Create(&questionModel).Error; err != nil {
		tsx.Rollback()
		return err
	}

	for _, option := range question.Options {

		var image, audio string
		isActive := true

		if option.Image != nil {
			imageData, err := file.UploadFile("option", option.Image)
			if err != nil {
				tsx.Rollback()

				return err
			}
			image = utils.GetURLObject(imageData.Url)

			err = tsx.Create(&models.File{
				Title:    imageData.Filename,
				Type:     imageData.FileType,
				Url:      image,
				IsActive: &isActive,
			}).Error

			if err != nil {
				tsx.Rollback()
				return err
			}
		}

		if option.Audio != nil {
			audioData, err := file.UploadFile("option", option.Audio)
			if err != nil {
				tsx.Rollback()
				return err
			}
			audio = utils.GetURLObject(audioData.Url)
			err = tsx.Create(&models.File{
				Title:    audioData.Filename,
				Type:     audioData.FileType,
				Url:      audio,
				IsActive: &isActive,
			}).Error

			if err != nil {
				tsx.Rollback()
				return err
			}
		}

		optionModel := models.Option{
			QuestionID: questionModel.ID,
			Image:      image,
			Audio:      audio,
			Text:       option.Text,
			IsCorrect:  option.IsCorrect,
		}

		if err := tsx.Create(&optionModel).Error; err != nil {
			tsx.Rollback()
			return err
		}
	}

	if question.NestedQuestionType != "CaseBased" {
		if question.QuestionSetID != 0 {
			var count int64
			// questionModel.QuestionSetID = &question.QuestionSetID

			err := tsx.Model(&models.QuestionSetQuestion{}).Where("question_set_id = ?", question.QuestionSetID).Count(&count).Error
			if err != nil {
				return err
			}

			err = tsx.Debug().Create(&models.QuestionSetQuestion{
				QuestionSetID: question.QuestionSetID,
				QuestionID:    questionModel.ID,
				Position:      uint(count) + 1,
			}).Error
			if err != nil {
				tsx.Rollback()
				return err
			}
		}
	}
	tsx.Commit()

	return nil
}
