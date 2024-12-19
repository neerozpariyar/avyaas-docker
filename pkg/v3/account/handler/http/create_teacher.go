package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
CreateTeacher is a Fiber handler function for processing requests to create a new teacher

Parameters:
  - handler: A pointer to the handler struct, typically representing the application's request
    handler. It is used to access the associated use case for creating a teacher.

Returns:
  - fiber.Handler: A Fiber handler function that can be registered to handle HTTP requests. The
    handler function processes requests, performs necessary validations, and delegates the teacher
    creation operation to the use case.
*/
func (handler *handler) CreateTeacher() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.TeacherCreateUpdateRequest

		// Parse the request json body to validate and extract the user information
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
		file, err := c.FormFile("image")
		if err != nil {
			requestBody.Image = nil
		} else {
			requestBody.Image = file
		}

		// Invoke the usecase for creation of teacher
		errMap = handler.usecase.CreateTeacher(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
