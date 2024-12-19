package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdateSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.SubjectCreateUpdateRequest
		// var subject *models.Subject

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

		// Change the SubjectID to uppercase
		requestBody.SubjectID = strings.ToUpper(requestBody.SubjectID)

		// Get the file from file field
		file, err := c.FormFile("file")
		if err != nil {
			requestBody.File = nil
		} else {
			requestBody.File = file
		}

		// Invoke request to the subject  usecase to update the subject
		errMap = handler.usecase.UpdateSubject(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
