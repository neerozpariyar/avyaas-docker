package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (r *Repository) UpdateTrueOrFalseQuestion(question *presenter.TypeQuestionPresenter) error {
	tsx := r.db.Begin()

	questionModel := models.TypeQuestion{
		Title: question.Title,
		// Type:         "TrueOrFalse", // Set the type to "TrueOrFalse"
		ForTest:      question.ForTest,
		SubjectID:    question.SubjectID,
		NegativeMark: question.NegativeMark,
		IsTrue:       question.IsTrue,
	}

	if err := tsx.Model(&models.TypeQuestion{}).Where("id = ?", question.ID).Updates(questionModel).Error; err != nil {
		tsx.Rollback()
		return err
	}
	tsx.Commit()
	return nil
}
