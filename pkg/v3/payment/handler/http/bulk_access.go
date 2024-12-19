package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) BulkAccessPaymentRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		errMap := make(map[string]string)

		// Parse the request body
		var request presenter.BulkAccessPaymentRequest
		err = c.BodyParser(&request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		request.File, err = c.FormFile("file")
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		extension := filepath.Ext(request.File.Filename)
		if extension != ".xlsx" {
			errMap["error"] = "Invalid file type. Please upload a .xlsx file."
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		err = validate.Struct(request)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Invoke request to the payment usecase to add the manual payment
		err = handler.usecase.BulkAcessPayment(&request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
