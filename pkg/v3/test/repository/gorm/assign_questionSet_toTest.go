package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) AssignQuestionSetToTest(testID, questionSetID uint) error {
	err := repo.db.Create(&models.TestQuestionSet{
		TestID:        testID,
		QuestionSetID: questionSetID,
	}).Error

	return err
}
