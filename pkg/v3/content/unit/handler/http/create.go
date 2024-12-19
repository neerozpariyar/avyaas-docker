package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateUnit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.UnitCreateUpdateRequest
		// var unit *models.Unit

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

		// Get the file from file field
		file, err := c.FormFile("file")
		if err != nil {
			requestBody.File = nil
		} else {
			requestBody.File = file
		}

		// Invoke request to the unit usecase to create the unit
		errMap = handler.usecase.CreateUnit(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
