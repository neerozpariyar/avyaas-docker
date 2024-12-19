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
func (repo *Repository) GetQuestionByID(id uint) (models.Question, error) {
	var question models.Question

	err := repo.db.Where("id = ?", id).First(&question).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Question{}, fmt.Errorf("question with id: '%d' not found", id)
		}

		return models.Question{}, err
	}

	return question, nil
}

func (repo *Repository) GetOptionByQuestionID(id uint) (*models.Option, error) {
	var option *models.Option

	err := repo.db.Where("question_id = ?", id).First(&option).Error
	if err != nil {
		return nil, err
	}

	return option, nil
}

func (repo *Repository) GetOptionsByQuestionID(id uint) ([]*models.Option, error) {
	var options []*models.Option

	err := repo.db.Where("question_id = ?", id).Find(&options).Error
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

func (repo *Repository) GetNestedQuestions(id uint) ([]presenter.QuestionListResponse, error) {
	var questions []models.Question

	err := repo.db.Where("case_question_id = ?", id).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	var result []presenter.QuestionListResponse

	for _, question := range questions {
		var options []models.Option

		err := repo.db.Where("question_id = ?", question.ID).Find(&options).Error
		if err != nil {
			return nil, err
		}

		var optionPresenters []presenter.OptionListPresenter

		for _, option := range options {
			var image, audio string

			if option.Image != "" {
				image = utils.GetFileURL(option.Image)
			}

			if option.Audio != "" {
				audio = utils.GetFileURL(option.Audio)
			}

			optionPresenters = append(optionPresenters, presenter.OptionListPresenter{
				ID:         option.ID,
				QuestionID: option.QuestionID,
				Image:      &image,
				Audio:      &audio,
				Text:       option.Text,
				IsCorrect:  &option.IsCorrect,
			})
		}

		result = append(result, presenter.QuestionListResponse{
			ID:           question.ID,
			Title:        question.Title,
			Description:  *question.Description,
			Type:         question.Type,
			ForTest:      question.ForTest,
			NegativeMark: question.NegativeMark,
			Options:      optionPresenters,
			// QuestionSetID: *question.QuestionSetID,
			IsTrue: question.IsTrue,
		})

	}

	return result, nil
}

func (repo *Repository) GetCorrectAnswersIDForQuestion(questionId uint) (uint, error) {

	var option uint

	err := repo.db.Select("id").Where("question_id = ? AND is_correct = true").Find(&option).Error

	if err != nil {
		return 0, err
	}

	return option, nil
}
