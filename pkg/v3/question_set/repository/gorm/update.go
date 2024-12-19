package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

/*
UpdateQuestionSet is a repository method responsible for updating the details of a question set in
the repository based on the provided question set's ID.

Parameters:
  - questionSet: A models.QuestionSet instance containing the updated details of the question set.

Returns:
  - err: An error, if any, encountered during the update operation.
*/
func (repo *Repository) UpdateQuestionSet(data presenter.CreateUpdateQuestionSetRequest) error {
	if data.File != nil {
		transaction := repo.db.Begin()

		fileData, err := file.UploadFile("question_set", data.File)
		if err != nil {
			return err
		}

		isActive := true
		urlObject := utils.GetURLObject(fileData.Url)

		err = transaction.Create(&models.File{
			Title:    fileData.Filename,
			Type:     fileData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		questionSet, err := repo.GetQuestionSetByID(data.ID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if questionSet.File != "" {
			var qsFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", questionSet.File).First(&qsFile).Error
			if err == nil {
				if err = repo.db.Model(models.File{}).Where("id = ?", qsFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		err = transaction.Where("id = ?", data.ID).Updates(&models.QuestionSet{
			Title:       data.Title,
			Description: data.Description,
			// TotalQuestions: data.TotalQuestions,
			Marks:    data.Marks,
			CourseID: data.CourseID,
			File:     urlObject,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		transaction.Commit()
		return err
	}

	return repo.db.Where("id = ?", data.ID).Updates(&models.QuestionSet{
		Title:       data.Title,
		Description: data.Description,
		// TotalQuestions: data.TotalQuestions,
		Marks:    data.Marks,
		CourseID: data.CourseID,
	}).Error
}
