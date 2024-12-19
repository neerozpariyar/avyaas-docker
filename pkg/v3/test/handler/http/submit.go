package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) SubmitTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.SubmitTestRequest

		testID := uint(utils.StringToUint(c.Params("id")))
		requestBody.TestID = testID

		// Parse the request json body to validate and extract the test information
		if err := c.BodyParser(&requestBody); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		// Validates the request validateBody using a validator and returns validation errors if present
		if err := validate.Struct(requestBody); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Get the logged in user id from Fiber's Locals to set in requestBody's CreatedBy
		requestBody.UserID = c.Locals("requester").(uint)

		// Invoke request to the test usecase to create the test
		if errMap = handler.usecase.SubmitTest(requestBody); len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
