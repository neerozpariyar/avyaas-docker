package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

/*
CreateQuestionSet is a repository function that creates the provided question set data in the database.

Parameters:
  - data: A pointer to the models.QuestionSet structure containing information about the question set
    to be created.

Returns:
  - error: An error, if any, encountered during the database operation. Returns nil on success.
*/
func (repo *Repository) CreateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) error {
	transaction := repo.db.Begin()

	var newQuestionSet *models.QuestionSet
	err := transaction.Create(&models.QuestionSet{
		Title:          data.Title,
		Description:    data.Description,
		TotalQuestions: data.TotalQuestions,
		Marks:          data.Marks,
		CourseID:       data.CourseID,
	}).Scan(&newQuestionSet).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	if data.File != nil {
		fileData, err := file.UploadFile("question_set", data.File)
		if err != nil {
			return err
		}

		var file *models.File
		isActive := true
		urlObject := utils.GetURLObject(fileData.Url)

		err = transaction.Create(&models.File{
			Title:    fileData.Filename,
			Type:     fileData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Scan(&file).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		newQuestionSet.File = file.Url
	}

	if len(data.Questions) > 0 {
		for _, question := range data.Questions {
			newQuestion, err := repo.questionRepo.CreateQuestion(question)
			if err != nil {
				transaction.Rollback()
				return err
			}

			err = transaction.Create(&models.QuestionSetQuestion{
				QuestionSetID: newQuestionSet.ID,
				QuestionID:    newQuestion.ID,
				Position:      question.Position,
			}).Error
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	transaction.Commit()
	return err
}
