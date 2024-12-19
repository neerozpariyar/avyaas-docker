package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"fmt"

	"gorm.io/gorm"
)

var transaction *gorm.DB

func (r *Repository) UpdateCaseQuestion(question presenter.CreateUpdateQuestionRequest) (*models.Question, error) {
	transaction = r.db.Begin()

	questionModel, err := r.updateQuestion(question)
	if err != nil {
		transaction.Rollback()
		return nil, err
	}

	err = r.updateNestedQuestions(question, questionModel)
	if err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return questionModel, nil
}

func (r *Repository) updateQuestion(question presenter.CreateUpdateQuestionRequest) (*models.Question, error) {
	var image, audio string
	isActive := true

	if question.Image != nil {
		imageData, err := file.UploadFile("question", question.Image)
		if err != nil {
			tsx.Rollback()

			return nil, err
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
			return nil, err
		}
	}

	if question.Audio != nil {
		audioData, err := file.UploadFile("question", question.Audio)
		if err != nil {
			tsx.Rollback()
			return nil, err
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
			return nil, err
		}
	}

	questionModel := models.Question{
		Title:          question.Title,
		Description:    question.Description,
		Image:          image,
		Audio:          audio,
		CaseQuestionID: nil,
		ForTest:        question.ForTest,
		SubjectID:      question.SubjectID,
		NegativeMark:   question.NegativeMark,
	}

	if err := transaction.Model(&questionModel).Where("id = ?", questionModel.ID).Updates(questionModel).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}
	return &questionModel, nil
}

func (r *Repository) updateNestedQuestions(question presenter.CreateUpdateQuestionRequest, questionModel *models.Question) error {
	nestQ, err := r.GetNestedQuestions(question.ID)
	if err != nil {
		return err
	}

	for i, nestedQuestion := range question.Questions {
		nestedQuestion.SubjectID = question.SubjectID
		nestedQuestion.CaseQuestionID = &questionModel.ID
		nestedQuestion.NestedQuestionType = questionModel.Type
		nestedQuestion.Options = question.Questions[i].Options
		nestedQuestion.ID = nestQ[i].ID
		typ, err := r.GetNestedQuestions(question.ID)
		if err != nil {
			return err
		}

		switch typ[i].Type {
		case "FillInTheBlanks":
			if err := r.UpdateFillInBlanksQuestion(&nestedQuestion); err != nil {
				return err
			}

		case "MCQ", "MultiAnswer":

			if err := r.UpdateMCQQuestion(&nestedQuestion); err != nil {
				return err
			}

		case "TrueOrFalse":

			if err := r.UpdateTrueOrFalseQuestion(&nestedQuestion); err != nil {
				return err
			}

		default:
			return fmt.Errorf("unknown nested question type: %s", question.NestedQuestionType)
		}
	}
	return nil
}
