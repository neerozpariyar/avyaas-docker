package gorm

import (
	"avyaas/internal/domain/presenter"
	"errors"
)

/*
UpdateQuestion is a repository function that updates the provided question data in the database.

Parameters:
  - question: A pointer to the models.Question struct containing information about the question to
    be updated.

Returns:
  - error: An error, if any, encountered during the database operation. Returns nil on success.
*/

func (repo *Repository) UpdateQuestion(data presenter.CreateUpdateQuestionRequest) error {
	switch data.Type {
	case "CaseBased":
		_, err := repo.UpdateCaseQuestion(data)
		if err != nil {
			return err
		}
	case "FillInTheBlanks":
		err := repo.UpdateFillInBlanksQuestion(&data)
		if err != nil {
			return err
		}
	case "MCQ", "MultiAnswer":

		err := repo.UpdateMCQQuestion(&data)
		if err != nil {
			return err
		}
	case "TrueFalse":
		err := repo.UpdateTrueOrFalseQuestion(&data)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid question type")
	}
	return nil
}

/*
func (repo *Repository) UpdateQuestion(data presenter.CreateUpdateQuestionRequest) error {
	// Begin a new transaction
	transaction := repo.db.Begin()

	question := &models.Question{
		Title:     data.Title,
		ForTest:   data.ForTest,
		SubjectID: data.SubjectID,
	}

	isActive := true

	// Check if a new image is provided
	if data.Image != nil {
		imageData, err := file.UploadFile("question", data.Image)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if question.Image != "" {
			err = utils.UpdateFileIsActive(question.Image, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		imageObject := utils.GetURLObject(imageData.Url)
		err = transaction.Create(&models.File{
			Title:    imageData.Filename,
			Type:     imageData.FileType,
			Url:      imageObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		question.Image = imageObject
	}

	// Check if a new audio is provided
	if data.Audio != nil {
		audioData, err := file.UploadFile("question", data.Audio)
		if err != nil {
			transaction.Rollback()
			return err
		}

		if question.Audio != "" {
			err = utils.UpdateFileIsActive(question.Audio, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		audioObject := utils.GetURLObject(audioData.Url)
		err = transaction.Create(&models.File{
			Title:    audioData.Filename,
			Type:     audioData.FileType,
			Url:      audioObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		question.Audio = audioObject
	}

	if err := transaction.Where("id = ?", data.ID).Updates(&question).Error; err != nil {
		transaction.Rollback()
		return err
	}

	if data.Answer != "A" && data.Answer != "a" && data.Answer != "B" && data.Answer != "b" && data.Answer != "C" && data.Answer != "c" && data.Answer != "D" && data.Answer != "d" {
		return fmt.Errorf("answer must be a valid option either A, B, C or D")
	}

	var option models.Option
	if err := transaction.Where("question_id = ?", data.ID).First(&option).Error; err != nil {
		transaction.Rollback()
		return err
	}

	option.Answer = strings.ToUpper(data.Answer)

	// Option A
	option.OptionA = data.Options[0].Text

	if data.Options[0].Image != nil {
		if option.ImageA != "" {
			err := utils.UpdateFileIsActive(option.ImageA, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		imageAData, err := file.UploadFile("option", data.Options[0].Image)
		if err != nil {
			transaction.Rollback()
			return err
		}

		imageAObject := utils.GetURLObject(imageAData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageAData.Filename,
			Type:     imageAData.FileType,
			Url:      imageAObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.ImageA = imageAObject
		option.OptionA = ""
	}

	if data.Options[0].Audio != nil {
		if option.AudioA != "" {
			err := utils.UpdateFileIsActive(option.AudioA, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		audioAData, err := file.UploadFile("option", data.Options[0].Audio)
		if err != nil {
			transaction.Rollback()
			return err
		}

		audioAObject := utils.GetURLObject(audioAData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioAData.Filename,
			Type:     audioAData.FileType,
			Url:      audioAObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.AudioA = audioAObject
	}

	// Option B
	option.OptionB = data.Options[1].Text

	if data.Options[1].Image != nil {
		if option.ImageB != "" {
			err := utils.UpdateFileIsActive(option.ImageB, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		imageBData, err := file.UploadFile("option", data.Options[1].Image)
		if err != nil {
			transaction.Rollback()
			return err
		}

		imageBObject := utils.GetURLObject(imageBData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageBData.Filename,
			Type:     imageBData.FileType,
			Url:      imageBObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.ImageB = imageBObject
	}

	if data.Options[1].Audio != nil {
		if option.AudioB != "" {
			err := utils.UpdateFileIsActive(option.AudioB, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		audioBData, err := file.UploadFile("option", data.Options[1].Audio)
		if err != nil {
			transaction.Rollback()
			return err
		}

		audioBObject := utils.GetURLObject(audioBData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioBData.Filename,
			Type:     audioBData.FileType,
			Url:      audioBObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.AudioB = audioBObject
	}

	// Option C
	option.OptionC = data.Options[2].Text

	if data.Options[2].Image != nil {
		if option.ImageC != "" {
			err := utils.UpdateFileIsActive(option.ImageC, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		imageCData, err := file.UploadFile("option", data.Options[2].Image)
		if err != nil {
			transaction.Rollback()
			return err
		}

		imageCObject := utils.GetURLObject(imageCData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageCData.Filename,
			Type:     imageCData.FileType,
			Url:      imageCObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.ImageC = imageCObject
	}

	if data.Options[2].Audio != nil {
		if option.AudioC != "" {
			err := utils.UpdateFileIsActive(option.AudioC, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		audioCData, err := file.UploadFile("option", data.Options[2].Audio)
		if err != nil {
			transaction.Rollback()
			return err
		}

		audioCObject := utils.GetURLObject(audioCData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioCData.Filename,
			Type:     audioCData.FileType,
			Url:      audioCObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.AudioC = audioCObject
	}

	// Option D
	option.OptionD = data.Options[3].Text

	if data.Options[3].Image != nil {
		if option.ImageD != "" {
			err := utils.UpdateFileIsActive(option.ImageD, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		imageDData, err := file.UploadFile("option", data.Options[3].Image)
		if err != nil {
			transaction.Rollback()
			return err
		}

		imageDObject := utils.GetURLObject(imageDData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageDData.Filename,
			Type:     imageDData.FileType,
			Url:      imageDObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.ImageD = imageDObject
	}

	if data.Options[3].Audio != nil {
		if option.AudioD != "" {
			err := utils.UpdateFileIsActive(option.AudioD, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		audioDData, err := file.UploadFile("option", data.Options[3].Audio)
		if err != nil {
			transaction.Rollback()
			return err
		}

		audioDObject := utils.GetURLObject(audioDData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioDData.Filename,
			Type:     audioDData.FileType,
			Url:      audioDObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return err
		}

		option.AudioD = audioDObject
	}

	if err := transaction.Where("id = ?", option.ID).Save(&option).Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
*/
