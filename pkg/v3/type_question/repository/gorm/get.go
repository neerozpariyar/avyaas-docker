package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
GetQuestionByID is a repository method responsible for retrieving a question from the database based
on its unique identifier (ID).

Parameters:
  - id: A uint representing the unique identifier (ID) of the question to be retrieved.

Returns:
  - question: A models.Question instance representing the retrieved question.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetTypeQuestionByID(id uint) (models.TypeQuestion, error) {
	var question models.TypeQuestion

	err := repo.db.Where("id = ?", id).First(&question).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.TypeQuestion{}, fmt.Errorf("question with id: '%d' not found", id)
		}

		return models.TypeQuestion{}, err
	}

	return question, nil
}

func (repo *Repository) GetTypeOptionByQuestionID(id uint) (*models.TypeOption, error) {
	var option *models.TypeOption

	err := repo.db.Where("question_id = ?", id).First(&option).Error
	if err != nil {
		return nil, err
	}

	return option, nil
}

func (repo *Repository) GetTypeOptionsByQuestionID(questionID uint) ([]*models.TypeOption, error) {
	var options []*models.TypeOption

	err := repo.db.Where("question_id = ?", questionID).Find(&options).Error
	if err != nil {
		return nil, err
	}

	return options, nil
}

func (repo *Repository) CheckIsBookmarked(userID, questionID uint) (bool, error) {
	var bookmark models.Bookmark

	// Retrieve the bookmark from the database based on given id
	err := repo.db.Select("id").Where("user_id = ? AND question_id = ?", userID, questionID).First(&bookmark).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (repo *Repository) GetNestedQuestions(id uint) ([]presenter.TypeQuestionListPresenter, error) {
	var questions []models.TypeQuestion

	err := repo.db.Where("case_question_id = ?", id).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	var result []presenter.TypeQuestionListPresenter

	for _, question := range questions {
		var options []models.TypeOption

		err := repo.db.Where("question_id = ?", question.ID).Find(&options).Error
		if err != nil {
			return nil, err
		}

		var optionPresenters []presenter.TypeOptionListPresenter

		for _, option := range options {
			var image, audio string

			if option.Image != nil && *option.Image != "" {
				image = utils.GetFileURL(*option.Image)
			}

			if option.Audio != nil && *option.Audio != "" {
				audio = utils.GetFileURL(*option.Audio)
			}

			optionPresenters = append(optionPresenters, presenter.TypeOptionListPresenter{
				ID:         option.ID,
				QuestionID: option.QuestionID,
				Image:      &image,
				Audio:      &audio,
				Text:       *option.Text,
				IsCorrect:  &option.IsCorrect,
			})
		}

		result = append(result, presenter.TypeQuestionListPresenter{
			ID:           question.ID,
			Title:        question.Title,
			Description:  question.Description,
			Type:         question.Type,
			ForTest:      question.ForTest,
			SubjectID:    question.SubjectID,
			NegativeMark: question.NegativeMark,
			Options:      optionPresenters,
			// QuestionSetID: *question.QuestionSetID,
			IsTrue: question.IsTrue,
		})

	}

	return result, nil
}
