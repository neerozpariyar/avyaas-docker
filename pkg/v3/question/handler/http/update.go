package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
UpdateQuestion handles the HTTP request to update a new question by sending a HTTP request to the
question usecase.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) UpdateQuestion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.CreateUpdateQuestionRequest
		// Convert and set the ID from the query parameter to requestBody
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))

		// Parse the request json body to validate and extract the question information

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		requestBody.Image, _ = c.FormFile("image")
		requestBody.Audio, _ = c.FormFile("audio")

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

		for i := 0; i < 4; i++ {
			title := c.FormValue(fmt.Sprintf("options[%d][title]", i))
			imageFile, _ := c.FormFile(fmt.Sprintf("options[%d][image]", i)) // Ignore error if file is not present
			audioFile, _ := c.FormFile(fmt.Sprintf("options[%d][audio]", i))

			// if file != nil && title != "" {
			// 	errMap["error"] = fmt.Sprintf("Both title and file can't coexist: %d", i)
			// 	return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			// }

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
				if fileType != "png" && fileType != "mpeg" && fileType != "mp3" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
					errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg,jpg & png and AUDIO type:mpeg & mp3 allowed", fileType).Error()
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

		// Initialize the translation function
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

		if errMap = handler.usecase.UpdateQuestion(requestBody); len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}

func (handler *handler) UpdateQuestionHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody = &presenter.CreateUpdateQuestionRequest{}

		requestBody.ID = uint(utils.StringToUint(c.Params("id")))

		// Common request body parsing and field assignments
		err := c.BodyParser(requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		requestBody, errMap = handler.ValidateAndMapQuestionOptions(requestBody, false, c)

		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Run validation on the final parsed request body
		validate, trans := utils.InitTranslator()
		if err := validate.Struct(requestBody); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		if errMap = handler.usecase.UpdateQuestion(*requestBody); len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
