package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) AssignQuestionSetToTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody presenter.AssignQuestionSetToTestRequest
		errMap := make(map[string]string)

		// Parse the request json body to validate and extract the question set information
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

		// Call the repository function to assign content to chapter
		if err := handler.usecase.AssignQuestionSetToTest(uint(requestBody.TestID), uint(requestBody.QuestionSetID)); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Return success response
		return c.JSON(presenter.SuccessResponse())
	}
}
