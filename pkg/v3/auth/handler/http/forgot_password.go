package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
ForgotPassword is a handler function for processing requests to initiate the password recovery process.
It expects a JSON request body containing the necessary information for initiating the forgot password
flow.

Parameters:
  - handler: A pointer to the handler struct, typically representing the application's request handler.
    It is used to access the associated use case for initiating the forgot password process.

Returns:
  - fiber.Handler: A Fiber handler function that can be registered to handle HTTP requests.
    The handler function processes requests, performs necessary validations,
    and delegates the forgot password operation to the use case.
*/
func (handler *handler) ForgotPassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.ForgotPasswordRequest

		// Parse the request json body to validate
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		// Validates the request validateBody using a validator and returns validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Call the usecase for password recovery
		err = handler.usecase.ForgotPassword(requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
