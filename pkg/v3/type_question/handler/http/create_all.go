package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateAllQuestion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		var requestBody presenter.TypeQuestionPresenter

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		image, err := c.FormFile("image")

		if err != nil {
			requestBody.Image = nil
		} else {
			requestBody.Image = image
		}

		audio, err := c.FormFile("audio")

		if err != nil {
			requestBody.Audio = nil
		} else {
			requestBody.Audio = audio
		}

		if requestBody.Image != nil {
			fileType := utils.GetFileType(requestBody.Image.Filename)
			if fileType != "png" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
				errMap["fileType"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg, jpg & png allowed", fileType).Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}

		}

		if requestBody.Audio != nil {
			fileType := utils.GetFileType(requestBody.Audio.Filename)
			if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
				errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only AUDIO type:mpeg & mp3 allowed", fileType).Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
		}

		switch requestBody.Type {

		case "MCQ", "TrueFalse", "FillInTheBlanks", "MultiAnswer":
			for i := 0; i < 4; i++ {

				var text string
				var correct bool
				var imageFile *multipart.FileHeader
				var audioFile *multipart.FileHeader
				correct, _ = strconv.ParseBool(c.FormValue(fmt.Sprintf("options.%d.isCorrect", i)))
				text = c.FormValue(fmt.Sprintf("options.%d.text", i))
				imageFile, _ = c.FormFile(fmt.Sprintf("options.%d.image", i))
				audioFile, _ = c.FormFile(fmt.Sprintf("options.%d.audio", i))
				var optionData presenter.TypeOptionPresenter

				if text != "" || imageFile != nil || audioFile != nil {
					optionData.Text = text

					optionData.IsCorrect = correct

					if imageFile != nil {
						fileType := utils.GetFileType(imageFile.Filename)
						if fileType != "png" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
							errMap["fileType"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg, jpg & png allowed", fileType).Error()
							return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
						}

						optionData.Image = imageFile
					}

					if audioFile != nil {
						fileType := utils.GetFileType(audioFile.Filename)
						if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
							errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only AUDIO type:mpeg & mp3 allowed", fileType).Error()
							return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
						}

						optionData.Audio = audioFile
					}
					requestBody.Options = append(requestBody.Options, optionData)
				}

			}

		case "CaseBased":
			return handler.CreateCaseBasedQuestion()(c)
		default:
			errMap["question_type"] = fmt.Sprintf("Invalid question type: %s", requestBody.Type)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		switch requestBody.Type {
		case "MCQ":
			correctCount := 0
			for _, option := range requestBody.Options {
				if option.IsCorrect {
					correctCount++
				}
			}
			if correctCount != 1 {
				errMap["error"] = "MCQ should have only one correct answer"
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}

		case "TrueFalse":

			if requestBody.IsTrue == nil {
				errMap["error"] = "TrueFalse question should have isTrue field"
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
		case "FillInTheBlanks":
			// Check if the option contains only text
			for _, option := range requestBody.Options {
				if option.Image != nil || option.Audio != nil {
					errMap["options"] = "FillInTheBlanks questions must have options of text only"
					return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
				}

			}
		case "MultiAnswer":
			correctCount := 0
			for i := range requestBody.Options {
				if requestBody.Options[i].IsCorrect {
					correctCount++
				}
			}
			if correctCount < 2 {
				errMap["error"] = "MultiAnswer should have atleast two correct answer"
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}

		}

		validate, trans := utils.InitTranslator()

		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		// if requestBody.Type != "TrueFalse" {
		// 	if len(requestBody.Options) < 1 {
		// 		errMap["error"] = fmt.Sprintf("Not enough options, provided options: %d", len(requestBody.Options))
		// 		return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		// 	}
		// }

		// Invoke the usecase for creation of question
		errMap = handler.usecase.CreateTypeQuestion(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}

//Insert this above bodyparser if it not all the field gets mapped

// requestBody.Title = c.FormValue("title")
// desc := c.FormValue("description")
// requestBody.Description = &desc
// requestBody.Type = c.FormValue("type")
// subID, _ := strconv.Atoi(c.FormValue("subjectID"))
// requestBody.SubjectID = uint(subID)
// forTest, _ := strconv.ParseBool(c.FormValue("forTest"))
// requestBody.ForTest = &forTest
// qID, _ := strconv.ParseUint(c.FormValue("questionSetID"), 10, 64)
// requestBody.QuestionSetID = uint(qID)
// requestBody.Image, _ = c.FormFile("image")
// requestBody.Audio, _ = c.FormFile("audio")

// if requestBody.Image != nil {
// 	fileType := utils.GetFileType(requestBody.Image.Filename)
// 	if fileType != "png" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
// 		errMap["fileType"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg, jpg & png allowed", fileType).Error()
// 		return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 	}

// }

// if requestBody.Audio != nil {
// 	fileType := utils.GetFileType(requestBody.Audio.Filename)
// 	if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
// 		errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only AUDIO type:mpeg & mp3 allowed", fileType).Error()
// 		return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 	}
// }
