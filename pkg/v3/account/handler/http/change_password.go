package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
ChangePassword is a handler function for processing requests to change a user's password.

Parameters:
  - handler: A pointer to the handler struct, typically representing the application's request handler.
    It is used to access the associated use case for changing passwords.

Returns:
  - fiber.Handler: A Fiber handler function that can be registered to handle HTTP requests.
    The handler function processes requests, performs necessary validations,
    and delegates the change password operation to the use case.
*/
func (handler *handler) ChangePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.ChangePasswordRequest

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

		// Set the userID from c.Local("requester") to requestBody
		requestBody.UserID = c.Locals("requester").(uint)

		// Call the usecase for change password
		errMap = handler.usecase.ChangePassword(requestBody)
		if len(errMap) != 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
