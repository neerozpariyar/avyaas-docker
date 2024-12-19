package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"errors"
)

func (repo *Repository) CreateQuestion(data presenter.CreateUpdateQuestionRequest) (question *models.Question, err error) {

	switch data.Type {
	case "CaseBased":
		question, err = repo.CreateCaseQuestion(data)
		if err != nil {
			return nil, err
		}
	case "FillInTheBlanks":
		err = repo.CreateFillInBlanksQuestion(&data)
		if err != nil {
			return nil, err
		}
	case "MCQ", "MultiAnswer":
		err := repo.CreateMCQQuestion(&data)
		if err != nil {
			return nil, err
		}
	case "TrueFalse":
		err := repo.CreateTrueOrFalseQuestion(&data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid question type")
	}

	return question, err
}

/*
CreateQuestion is a repository function that inserts the provided question data in the database.

Parameters:
  - data: A pointer to the models.Question structure containing information about the question to
    be created.

Returns:
  - error: An error, if any, encountered during the database operation. Returns nil on success.
*/

//Old code
/*
func (repo *Repository) CreateQuestion(data presenter.CreateUpdateQuestionRequest) (*models.Question, error) {
	// Begin a new transaction
	transaction := repo.db.Begin()
	// var correctOption string

	question := &models.Question{
		Title:        data.Title,
		ForTest:      data.ForTest,
		SubjectID:    data.SubjectID,
		NegativeMark: data.NegativeMark,
	}

	isActive := true

	if data.Image != nil {
		// Upload the file for the question's image
		imageData, err := file.UploadFile("question", data.Image)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		urlObject := utils.GetURLObject(imageData.Url)
		err = transaction.Create(&models.File{
			Title:    imageData.Filename,
			Type:     imageData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		question.Image = urlObject
	}

	if data.Audio != nil {
		// Upload the file for the question's audio
		audioData, err := file.UploadFile("question", data.Audio)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		urlObject := utils.GetURLObject(audioData.Url)
		err = transaction.Create(&models.File{
			Title:    audioData.Filename,
			Type:     audioData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		question.Image = urlObject
	}

	if err := transaction.Create(&question).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	if data.Answer != "A" && data.Answer != "a" && data.Answer != "B" && data.Answer != "b" && data.Answer != "C" && data.Answer != "c" && data.Answer != "D" && data.Answer != "d" {
		return nil, fmt.Errorf("answer must be a valid option either A, B, C or D")
	}

	option := models.Option{
		QuestionID: question.ID,
		Answer:     strings.ToUpper(data.Answer), // Set the answer as the correct option
	}

	// Option A
	option.OptionA = data.Options[0].Text
	if data.Options[0].Image != nil {
		imageAData, err := file.UploadFile("option", data.Options[0].Image)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		imageAObject := utils.GetURLObject(imageAData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageAData.Filename,
			Type:     imageAData.FileType,
			Url:      imageAObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.ImageA = imageAObject
	}

	if data.Options[0].Audio != nil {
		audioAData, err := file.UploadFile("option", data.Options[0].Audio)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		audioAObject := utils.GetURLObject(audioAData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioAData.Filename,
			Type:     audioAData.FileType,
			Url:      audioAObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.AudioA = audioAObject
	}

	// Option B
	option.OptionB = data.Options[1].Text
	if data.Options[1].Image != nil {
		imageBData, err := file.UploadFile("option", data.Options[1].Image)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		imageBObject := utils.GetURLObject(imageBData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageBData.Filename,
			Type:     imageBData.FileType,
			Url:      imageBObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.ImageB = imageBObject
	}

	if data.Options[1].Audio != nil {
		audioBData, err := file.UploadFile("option", data.Options[1].Audio)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		audioBObject := utils.GetURLObject(audioBData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioBData.Filename,
			Type:     audioBData.FileType,
			Url:      audioBObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.AudioB = audioBObject
	}

	// Option C
	option.OptionC = data.Options[2].Text
	if data.Options[2].Image != nil {
		imageCData, err := file.UploadFile("option", data.Options[2].Image)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		imageCObject := utils.GetURLObject(imageCData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageCData.Filename,
			Type:     imageCData.FileType,
			Url:      imageCObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.ImageC = imageCObject
	}

	if data.Options[2].Audio != nil {
		audioCData, err := file.UploadFile("option", data.Options[2].Audio)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		audioCObject := utils.GetURLObject(audioCData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioCData.Filename,
			Type:     audioCData.FileType,
			Url:      audioCObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.AudioC = audioCObject
	}

	// Option D
	option.OptionD = data.Options[3].Text
	if data.Options[3].Image != nil {
		imageDData, err := file.UploadFile("option", data.Options[3].Image)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		imageDObject := utils.GetURLObject(imageDData.Url)

		if err := transaction.Create(&models.File{
			Title:    imageDData.Filename,
			Type:     imageDData.FileType,
			Url:      imageDObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.ImageD = imageDObject
	}

	if data.Options[3].Audio != nil {
		audioDData, err := file.UploadFile("option", data.Options[3].Audio)
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		audioDObject := utils.GetURLObject(audioDData.Url)

		if err := transaction.Create(&models.File{
			Title:    audioDData.Filename,
			Type:     audioDData.FileType,
			Url:      audioDObject,
			IsActive: &isActive,
		}).Error; err != nil {
			transaction.Rollback()
			return nil, err
		}

		option.AudioD = audioDObject
	}

	if err := transaction.Create(&option).Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	if data.QuestionSetID != 0 {
		var count int64
		err := repo.db.Model(&models.QuestionSetQuestion{}).Where("question_set_id = ?", data.QuestionSetID).Count(&count).Error
		if err != nil {
			transaction.Rollback()
			return nil, err
		}

		err = transaction.Create(&models.QuestionSetQuestion{
			QuestionSetID: data.QuestionSetID,
			QuestionID:    question.ID,
			Position:      uint(count) + 1,
		}).Error
		if err != nil {
			transaction.Rollback()
			return nil, err
		}
	}

	transaction.Commit()
	return question, nil
}
*/
// Check if any file field is not nil
// option.OptionA = data.Options[0].Title
// if data.Options[0].Title != "" {
// 	option.OptionA = data.Options[0].Title

// 	data.Options[0].File = nil
// } else if data.Options[0].File != nil {
// 	fileAData, err := file.UploadFile("option", data.Options[0].File)
// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}

// 	urlAObject := utils.GetURLObject(fileAData.Url)

// 	if err := transaction.Create(&models.File{
// 		Title:    fileAData.Filename,
// 		Type:     fileAData.FileType,
// 		Url:      urlAObject,
// 		IsActive: &isActive,
// 	}).Error; err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}
// 	option.UrlA = urlAObject
// 	option.FileType = fileAData.FileType
// }

// if data.Options[1].Title != "" {
// 	option.OptionB = data.Options[1].Title
// 	data.Options[1].File = nil
// } else if data.Options[1].File != nil {
// 	fileBData, err := file.UploadFile("option", data.Options[0].File)
// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}

// 	urlBObject := utils.GetURLObject(fileBData.Url)
// 	if err := transaction.Create(&models.File{
// 		Title:    fileBData.Filename,
// 		Type:     fileBData.FileType,
// 		Url:      urlBObject,
// 		IsActive: &isActive,
// 	}).Error; err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}
// 	option.UrlB = urlBObject
// 	option.FileType = fileBData.FileType
// }

// if data.Options[2].Title != "" {
// 	option.OptionC = data.Options[2].Title
// 	data.Options[2].File = nil
// } else if data.Options[2].File != nil {
// 	fileCData, err := file.UploadFile("option", data.Options[2].File)
// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}

// 	urlCObject := utils.GetURLObject(fileCData.Url)
// 	if err := transaction.Create(&models.File{
// 		Title:    fileCData.Filename,
// 		Type:     fileCData.FileType,
// 		Url:      urlCObject,
// 		IsActive: &isActive,
// 	}).Error; err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}
// 	option.UrlC = urlCObject
// 	option.FileType = fileCData.FileType
// }

// if data.Options[3].Title != "" {
// 	option.OptionD = data.Options[3].Title

// 	data.Options[3].File = nil
// } else if data.Options[3].File != nil {
// 	fileDData, err := file.UploadFile("option", data.Options[3].File)
// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}

// 	urlDObject := utils.GetURLObject(fileDData.Url)
// 	if err := transaction.Create(&models.File{
// 		Title:    fileDData.Filename,
// 		Type:     fileDData.FileType,
// 		Url:      urlDObject,
// 		IsActive: &isActive,
// 	}).Error; err != nil {
// 		transaction.Rollback()
// 		return nil, err
// 	}
// 	option.UrlD = urlDObject
// 	option.FileType = fileDData.FileType
// }

// } else {
// option := models.Option{
// 	QuestionID: question.ID,
// 	OptionA:    data.OptionA,
// 	OptionB:    data.OptionB,
// 	OptionC:    data.OptionC,
// 	OptionD:    data.OptionD,
// 	Answer:     data.Answer, // Set the correct option URL
// }
// if err := transaction.Create(&option).Error; err != nil {
// 	transaction.Rollback()
// 	return nil, err
// }
// }

///////////////////////////??????????????????????????????????////////////////////
// func (repo *repository) CreateQuestion(data presenter.CreateQuestionRequest) error {
// 	// Begin a new transaction
// 	transaction := repo.db.Begin()

// 	question := &models.Question{
// 		Text: data.Text,
// 	}

// 	var options []models.Option

// 	for i, option := range data.Options {
// 		fileData, err := file.UploadFile("course_group", data.OptionUrl[i])
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		isActive := true
// 		err = transaction.Create(&models.File{
// 			Name:     fileData.Filename,
// 			Type:     fileData.FileType,
// 			Url:      fileData.Url,
// 			IsActive: &isActive,
// 		}).Error

// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		op := models.Option{
// 			Text: option,
// 			Url:  fileData.Url,
// 		}
// 		options = append(options, op)
// 	}

// 	question.Options = options

// 	if err := transaction.Create(question).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	transaction.Commit()

// 	return nil
// }
/////////**************************************???????????///////////////////

// func (repo *repository) CreateQuestion(data presenter.CreateQuestionRequest) error {
// 	transaction := repo.db.Begin()

// 	question := &models.Question{
// 		Text: data.Text,
// 	}
// 	var options []models.Option
// 	for _, optionData := range data.Options {
// 		op := models.Option{
// 			Text: optionData.Text,
// 			// Assuming you want to store the file URL, adjust accordingly
// 			Url: optionData.Url.Filename, // Use the appropriate field from FileHeader
// 		}
// 		options = append(options, op)
// 	}

// 	question.Options = options

// 	if err := transaction.Create(question).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	transaction.Commit()
// 	return nil
// }

// func (repo *repository) CreateQuestion(data presenter.CreateUpdateQuestionRequest) error {
// 	transaction := repo.db.Begin()

// 	// Upload file for the main question and save file details
// 	fileData, err := file.UploadFile("course_group", data.Image)
// 	if err != nil {
// 		return err
// 	}

// 	isActive := true
// 	if err := transaction.Create(&models.File{
// 		Name:     fileData.Filename,
// 		Type:     fileData.FileType,
// 		Url:      fileData.Url,
// 		IsActive: &isActive,
// 	}).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	// Create a new question
// 	question := &models.Question{
// 		Name:          data.Name,
// 		Description:   data.Description,
// 		Image:         fileData.Url,
// 		Answer:        data.Answer,
// 		Position:      data.Position,
// 		ForTest:       data.ForTest,
// 		SubjectID:     data.SubjectID,
// 		QuestionSetID: data.QuestionSetID,
// 	}

// 	// Create options
// 	var options []models.Option
// 	for _, optionTitle := range data.Options {
// 		// Extract file from the data for each option
// 		optionFile, err := c.FormFile("option_file")
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		// Upload file for each option and save file details
// 		optionFileData, err := file.UploadFile("course_group", optionFile)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		option := models.Option{
// 			Title:    optionTitle,
// 			Option:   optionFileData.Url,
// 			FileType: optionFileData.FileType,
// 		}
// 		options = append(options, option)
// 	}

// 	question.Options = options

// 	if err := transaction.Create(question).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	transaction.Commit()
// 	return nil
// }
