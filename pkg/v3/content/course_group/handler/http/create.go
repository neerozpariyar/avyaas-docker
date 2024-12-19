package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
CreateCourseGroup handles the HTTP request to create a new course group by sending a HTTP request to the
course group usecase.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) CreateCourseGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.CourseGroupCreateUpdateRequest

		// Parse the request json body to validate and extract the course group information
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

		// Get the file from file field
		file, err := c.FormFile("file")
		if err != nil {
			requestBody.File = nil
		} else {
			requestBody.File = file
		}

		// Change the GroupID to uppercase
		requestBody.GroupID = strings.ToUpper(requestBody.GroupID)

		// Invoke request to the course group usecase to create the course group
		errMap = handler.usecase.CreateCourseGroup(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
