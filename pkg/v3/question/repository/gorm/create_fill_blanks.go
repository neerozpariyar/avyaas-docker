package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (r *Repository) CreateFillInBlanksQuestion(question *presenter.CreateUpdateQuestionRequest) error {
	tsx := r.db.Begin()

	questionModel := models.Question{
		Timestamp:    models.Timestamp{ID: question.ID},
		Title:        question.Title, //paragraph in title
		Description:  question.Description,
		Type:         "FillInTheBlanks",
		ForTest:      question.ForTest,
		SubjectID:    question.SubjectID,
		NegativeMark: question.NegativeMark,
	}

	if question.NestedQuestionType == "CaseBased" {
		questionModel.CaseQuestionID = question.CaseQuestionID

	}

	fmt.Printf("len(question.Options): %v\n", len(question.Options))
	// Extract text from each option and add to questionModel.Options
	for _, option := range question.Options {

		questionModel.Options = append(questionModel.Options, models.Option{
			Text:  option.Text,
			Image: "",
			Audio: "",
		})
	}

	if err := tsx.Create(&questionModel).Error; err != nil {
		return err
	}

	if question.NestedQuestionType != "CaseBased" {

		if question.QuestionSetID != 0 {
			var count int64

			// questionModel.QuestionSetID = &question.QuestionSetID

			err := tsx.Model(&models.QuestionSetQuestion{}).Where("question_set_id = ?", question.QuestionSetID).Count(&count).Error
			if err != nil {
				tsx.Rollback()
				return err
			}

			err = tsx.Create(&models.QuestionSetQuestion{
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
