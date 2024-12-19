package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
GetQuestionSetByID is a repository method responsible for retrieving a question set from the database
based on its unique identifier (ID).

Parameters:
  - id: A uint representing the unique identifier (ID) of the question set to be retrieved.

Returns:
  - questionSet: A models.QuestionSet instance representing the retrieved question set.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetQuestionSetByID(id uint) (models.QuestionSet, error) {
	var questionSet models.QuestionSet

	// Retrieve the question set from the database based on given id
	err := repo.db.Where("id = ?", id).Preload("Questions").First(&questionSet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.QuestionSet{}, fmt.Errorf("question set with id: '%d' not found", id)
		}

		return models.QuestionSet{}, err
	}

	return questionSet, nil
}

func (repo *Repository) GetQuestionSetByTitleAndCourseID(title string, courseID uint) (models.QuestionSet, error) {
	var questionSet models.QuestionSet

	// Retrieve the question set from the database based on given id
	err := repo.db.Where("title = ? AND course_id = ?", title, courseID).First(&questionSet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.QuestionSet{}, fmt.Errorf("question set with title: '%s' not found", title)
		}

		return models.QuestionSet{}, err
	}

	return questionSet, nil
}

func (repo *Repository) GetQuestionSetQuestion(questionSetID, questionID uint) (*models.QuestionSetQuestion, error) {
	var qsQuestion *models.QuestionSetQuestion

	err := repo.db.Where("question_set_id = ? AND question_id = ?", questionSetID, questionID).First(&qsQuestion).Error
	if err != nil {
		return nil, err
	}

	return qsQuestion, nil
}
