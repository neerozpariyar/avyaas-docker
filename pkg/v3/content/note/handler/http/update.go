package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdateNote() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.NoteCreateUpdateRequest

		// Parse and validate the request json body
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		// Validate the request validateBody using a validator and return the translated validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Convert and set the ID from the route parameter to requestBody
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))

		// Get the file from file field
		file, _ := c.FormFile("file")
		if file != nil {
			requestBody.File = file
		}

		// Invoke request to the note  usecase to update the note
		errMap = handler.usecase.UpdateNote(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}