package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
)

func (r *Repository) UpdateMCQQuestion(question *presenter.CreateUpdateQuestionRequest) error {

	tsx := r.db.Begin()
	var image, audio string
	isActive := true

	if question.Image != nil {
		imageData, err := file.UploadFile("question", question.Image)
		if err != nil {
			tsx.Rollback()

			return err
		}
		image = utils.GetURLObject(imageData.Url)

		err = tsx.Create(&models.File{
			Title:    imageData.Filename,
			Type:     imageData.FileType,
			Url:      image,
			IsActive: &isActive,
		}).Error

		if err != nil {
			tsx.Rollback()
			return err
		}
	}

	if question.Audio != nil {
		audioData, err := file.UploadFile("question", question.Audio)
		if err != nil {
			tsx.Rollback()
			return err
		}
		audio = utils.GetURLObject(audioData.Url)
		err = tsx.Create(&models.File{
			Title:    audioData.Filename,
			Type:     audioData.FileType,
			Url:      audio,
			IsActive: &isActive,
		}).Error

		if err != nil {
			tsx.Rollback()
			return err
		}
	}
	questionModel := models.Question{
		Timestamp:    models.Timestamp{ID: question.ID},
		Title:        question.Title,
		Image:        image,
		Audio:        audio,
		ForTest:      question.ForTest,
		SubjectID:    question.SubjectID,
		NegativeMark: question.NegativeMark,
	}

	if err := r.db.Model(&models.TypeQuestion{}).Where("id = ?", question.ID).Updates(questionModel).Error; err != nil {
		tsx.Rollback()
		return err
	}
	opts, err := r.GetOptionsByQuestionID(question.ID)
	if err != nil {
		return err
	}

	// Update each option in the database
	for i, option := range opts {
		var image, audio string

		if option.Image != "" {
			if question.Options[i].Image != nil {
				imageData, err := file.UploadFile("option", question.Options[i].Image)
				if err != nil {
					tsx.Rollback()
					return err
				}

				image = utils.GetURLObject(imageData.Url)
			}
		}

		if option.Audio != "" {
			if question.Options[i].Audio != nil {

				audioData, err := file.UploadFile("option", question.Options[i].Audio)
				if err != nil {
					tsx.Rollback()
					return err
				}
				audio = utils.GetURLObject(audioData.Url)
			}
		}

		// Create a map for the update as updates did not send false value for update
		updateMap := map[string]interface{}{
			"Image":     &image,
			"Audio":     &audio,
			"Text":      &question.Options[i].Text,
			"IsCorrect": question.Options[i].IsCorrect,
		}
		if err := tsx.Model(&models.TypeOption{}).Where("id = ?", option.ID).Updates(updateMap).Error; err != nil {
			tsx.Rollback()
			return err
		}
	}
	tsx.Commit()

	return nil
}
