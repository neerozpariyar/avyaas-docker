package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"fmt"

	"gorm.io/gorm"
)

// var question presenter.CreateUpdateQuestionRequest

// questionModel=parent question
// question=input from handler
// nestedQuestionModel=child question
var tsx *gorm.DB

func (r *Repository) CreateCaseQuestion(question presenter.CreateUpdateQuestionRequest) (*models.Question, error) {
	tsx = r.db.Begin()

	questionModel, err := r.createQuestion(question)
	if err != nil {
		tsx.Rollback()
		return nil, err
	}
	err = r.createNestedQuestions(question, questionModel)
	if err != nil {
		tsx.Rollback()
		return nil, err
	}

	if question.QuestionSetID != 0 {
		var count int64
		err := tsx.Model(&models.QuestionSetQuestion{}).Where("question_set_id = ?", question.QuestionSetID).Count(&count).Error
		if err != nil {
			tsx.Rollback()
			return nil, err
		}

		err = tsx.Create(&models.QuestionSetQuestion{
			QuestionSetID: question.QuestionSetID,
			QuestionID:    questionModel.ID,
			Position:      uint(count) + 1,
		}).Error
		if err != nil {
			tsx.Rollback()
			return nil, err
		}
	}
	tsx.Commit()
	return questionModel, nil
}

func (r *Repository) createQuestion(question presenter.CreateUpdateQuestionRequest) (*models.Question, error) {
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
			tsx.Rollback()
			return nil, err
		}
	}
	questionModel := models.Question{
		Timestamp:      models.Timestamp{ID: question.ID},
		Title:          question.Title,
		Image:          image,
		Audio:          audio,
		Description:    question.Description,
		Type:           question.Type,
		CaseQuestionID: nil,
		ForTest:        question.ForTest,
		SubjectID:      question.SubjectID,
		NegativeMark:   question.NegativeMark,
		// QuestionSetID:  &question.QuestionSetID,
	}
	if err := tsx.Create(&questionModel).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}
	return &questionModel, nil
}

func (r *Repository) createNestedQuestions(question presenter.CreateUpdateQuestionRequest, questionModel *models.Question) error {

	for i, nestedQuestion := range question.Questions {
		nestedQuestion.SubjectID = question.SubjectID

		nestedQuestion.CaseQuestionID = &questionModel.ID
		nestedQuestion.NestedQuestionType = questionModel.Type
		nestedQuestion.Options = question.Questions[i].Options

		switch question.Questions[i].Type {
		case "FillInTheBlanks":

			if err := r.CreateFillInBlanksQuestion(&nestedQuestion); err != nil {
				return err
			}
		case "MCQ", "MultiAnswer":

			if err := r.CreateMCQQuestion(&nestedQuestion); err != nil {
				return err
			}
		case "TrueOrFalse":

			if err := r.CreateTrueOrFalseQuestion(&nestedQuestion); err != nil {
				return err
			}
		default:

			return fmt.Errorf("unknown nested question type: %s", question.NestedQuestionType)
		}

	}
	return nil
}
