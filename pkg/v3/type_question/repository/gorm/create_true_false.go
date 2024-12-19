package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (r *Repository) CreateTrueOrFalseQuestion(question *presenter.TypeQuestionPresenter) error {
	tsx := r.db.Begin()

	questionModel := models.TypeQuestion{
		Timestamp:    models.Timestamp{ID: question.ID},
		Title:        question.Title,
		Type:         "TrueOrFalse", // Set the type to "TrueOrFalse"
		ForTest:      question.ForTest,
		SubjectID:    question.SubjectID,
		NegativeMark: question.NegativeMark,
		IsTrue:       question.IsTrue,
	}
	if question.NestedQuestionType == "CaseBased" {
		questionModel.CaseQuestionID = question.CaseQuestionID
	}
	if err := tsx.Create(&questionModel).Error; err != nil {
		return err
	}
	if question.NestedQuestionType != "CaseBased" {

		if question.QuestionSetID != 0 {

			// questionModel.QuestionSetID = &question.QuestionSetID

			var count int64

			err := tsx.Model(&models.QuestionSetQuestion{}).Where("question_set_id = ?", question.QuestionSetID).Count(&count).Error
			if err != nil {
				tsx.Rollback()
				return err
			}

			err = tsx.Create(&models.QuestionSetQuestion{
				QuestionSetID:  question.QuestionSetID,
				TypeQuestionID: questionModel.ID,
				Position:       uint(count) + 1,
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
