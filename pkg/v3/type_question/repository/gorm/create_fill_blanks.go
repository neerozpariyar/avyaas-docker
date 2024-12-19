package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (r *Repository) CreateFillInBlanksQuestion(question *presenter.TypeQuestionPresenter) error {
	tsx := r.db.Begin()
	// var image, audio string
	// isActive := true

	// if question.Image != nil {
	// 	imageData, err := file.UploadFile("question", question.Image)
	// 	if err != nil {
	// 		tsx.Rollback()

	// 		return err
	// 	}
	// 	image = utils.GetURLObject(imageData.Url)

	// 	err = tsx.Create(&models.File{
	// 		Title:    imageData.Filename,
	// 		Type:     imageData.FileType,
	// 		Url:      image,
	// 		IsActive: &isActive,
	// 	}).Error

	// 	if err != nil {
	// 		tsx.Rollback()
	// 		return err
	// 	}
	// }

	// if question.Audio != nil {
	// 	audioData, err := file.UploadFile("question", question.Audio)
	// 	if err != nil {
	// 		tsx.Rollback()
	// 		return err
	// 	}
	// 	audio = utils.GetURLObject(audioData.Url)
	// 	err = tsx.Create(&models.File{
	// 		Title:    audioData.Filename,
	// 		Type:     audioData.FileType,
	// 		Url:      audio,
	// 		IsActive: &isActive,
	// 	}).Error

	// 	if err != nil {
	// 		tsx.Rollback()
	// 		return err
	// 	}
	// }
	questionModel := models.TypeQuestion{
		Timestamp: models.Timestamp{ID: question.ID},
		Title:     question.Title, //paragraph in title
		// Image:        image,
		// Audio:        audio,
		Description:  nil,
		Type:         "FillInTheBlanks",
		ForTest:      question.ForTest,
		SubjectID:    question.SubjectID,
		NegativeMark: question.NegativeMark,
	}

	if question.NestedQuestionType == "CaseBased" {
		questionModel.CaseQuestionID = question.CaseQuestionID

	}

	// Extract text from each option and add to questionModel.Options
	for _, option := range question.Options {

		questionModel.Options = append(questionModel.Options, models.TypeOption{
			Text:  &option.Text,
			Image: nil,
			Audio: nil,
		})
	}

	if err := tsx.Create(&questionModel).Error; err != nil {
		return err
	}

	if question.NestedQuestionType != "CaseBased" {

		if question.QuestionSetID != 0 {
			var count int64

			// questionModel.QuestionSetID = &question.QuestionSetID

			err := tsx.Model(&models.QuestionSetQuestion{}).Where("question_set_id = ?", question.QuestionSetID).Count(&count).Error
			if err != nil {
				tsx.Rollback()
				return err
			}

			err = tsx.Create(&models.QuestionSetQuestion{
				QuestionSetID:  question.QuestionSetID,
				TypeQuestionID: questionModel.ID,
				Position:       uint(count) + 1,
			}).Error
			if err != nil {
				tsx.Rollback()
				return err
			}
		}
	}

	tsx.Commit()

	return nil
}
