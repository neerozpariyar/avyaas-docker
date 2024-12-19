package http

import (
	"avyaas/internal/domain/presenter"

	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateQuestion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Unified request body structure to handle both single and type-based question creation
		var requestBody = &presenter.CreateUpdateQuestionRequest{}

		// Parse basic request body
		err := c.BodyParser(requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		requestBody, errMap = handler.ValidateAndMapQuestionOptions(requestBody, true, c)

		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Invoke the correct usecase based on type

		errMap = handler.usecase.CreateQuestion(*requestBody)

		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}

//THIS IS THE OLD CODE FOR CREATING QUESTION
/*
CreateQuestion handles the HTTP request to create a new question by sending a HTTP request to the
question usecase.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.

func (handler *handler) CreateQuestion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.CreateUpdateQuestionRequest
		// Parse options

		for i := 0; i < 4; i++ {
			title := c.FormValue(fmt.Sprintf("options[%d][title]", i))
			imageFile, _ := c.FormFile(fmt.Sprintf("options[%d][image]", i)) // Ignore error if file is not present
			audioFile, _ := c.FormFile(fmt.Sprintf("options[%d][audio]", i))

			var optionData presenter.OptionCreate

			optionData.Text = title

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
				fmt.Printf("fileType: %v\n", fileType)
				if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
					errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only AUDIO type:mpeg & mp3 allowed", fileType).Error()
					return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
				}

				optionData.Audio = audioFile
			}

			requestBody.Options = append(requestBody.Options, optionData)
		}

		if len(requestBody.Options) < 4 {
			errMap["error"] = fmt.Sprintf("Not enough options, provided options: %d", len(requestBody.Options))
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		validate, trans := utils.InitTranslator()

		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Get the file from the "image" field
		qImage, _ := c.FormFile("image")
		if qImage != nil {
			fileType := utils.GetFileType(qImage.Filename)
			if fileType != "png" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
				errMap["image"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg, jpg & png for Image allowed", fileType).Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}

			requestBody.Image = qImage
		}

		// Get the file from the "audio" field
		qAudio, _ := c.FormFile("audio")
		if qAudio != nil {
			fileType := utils.GetFileType(qAudio.Filename)
			if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
				errMap["audio"] = fmt.Errorf("file type of %v not allowed: only AUDIO type: mpeg & mp3 for Audio allowed", fileType).Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}

			requestBody.Audio = qAudio
		}
		fmt.Printf("requestBody: %v\n", requestBody)
		// Invoke the usecase for creation of question
		errMap = handler.usecase.CreateQuestion(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
*/
