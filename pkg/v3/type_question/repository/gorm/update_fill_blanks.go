package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (r *Repository) UpdateFillInBlanksQuestion(question *presenter.TypeQuestionPresenter) error {

	tsx := r.db.Begin()

	questionModel := models.TypeQuestion{
		Timestamp: models.Timestamp{ID: question.ID},
		Title:     question.Title,
		// Type:         "FillInTheBlanks",
		ForTest:      question.ForTest,
		SubjectID:    question.SubjectID,
		NegativeMark: question.NegativeMark,
	}

	if err := tsx.Model(&models.TypeQuestion{}).Where("id = ?", question.ID).Updates(questionModel).Error; err != nil {
		tsx.Rollback()
		return err
	}

	opts, err := r.GetTypeOptionsByQuestionID(question.ID)
	if err != nil {
		return err
	}

	// Extract text from each option and add to questionModel.Options
	for i, option := range opts {

		// Update the ID of the option in question.Options with the ID of the option from the database.
		question.Options[i].ID = option.ID
		optionModel := models.TypeOption{
			Text:  &question.Options[i].Text,
			Image: nil,
			Audio: nil,
		}
		if err := tsx.Model(&models.TypeOption{}).Where("id = ?", option.ID).Updates(optionModel).Error; err != nil {
			tsx.Rollback()
			return err
		}
	}

	tsx.Commit()

	return nil
}
