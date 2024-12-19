package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) AssignQuestionsToQuestionSet(questionSetID uint, questionIDs []uint) error {
	var qsQuestions []models.QuestionSetQuestion

	err := repo.db.Where("question_set_id = ?", questionSetID).Order("position").Find(&qsQuestions).Error
	if err != nil {
		return err
	}

	position := 0
	if len(qsQuestions) == 0 {
		position = 1
	} else {
		position = int(qsQuestions[len(qsQuestions)-1].Position) + 1
	}

	transaction := repo.db.Begin()

	for _, qID := range questionIDs {
		var question models.Question
		var questionSet models.QuestionSet

		err := repo.db.Model(&models.QuestionSet{}).Where("id = ?", questionSetID).First(&questionSet).Error
		if err != nil {
			return err
		}

		err = repo.db.Model(&models.Question{}).Where("id = ?", qID).First(&question).Error

		if err != nil {
			return err
		}

		err = transaction.Model(&questionSet).Association("Questions").Append(&question)

		if err != nil {
			return err
		}

		// err = transaction.Create(&models.QuestionSetQuestion{
		// 	QuestionSetID: questionSetID,
		// 	QuestionID:    qID,
		// 	Position:      uint(position),
		// }).Error
		// if err != nil {
		// 	transaction.Rollback()
		// 	return err
		// }

		position += 1
	}

	transaction.Commit()
	return nil
}
