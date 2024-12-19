package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// update profile by self only
func (handler *handler) UpdateStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.StudentCreateUpdateRequest

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

		// Retrieve the student id from Fiber context's params
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))
		loggedInUser := c.Locals("requester").(uint)

		// Get the file from file field
		file, err := c.FormFile("image")
		if err != nil {
			requestBody.Image = nil
		} else {
			requestBody.Image = file
		}

		// Invoke the update of student
		if loggedInUser == requestBody.ID {
			errMap = handler.usecase.UpdateStudent(requestBody)

		} else {
			errMap["user_access"] = "not authorised to update with this access level"
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))

		}
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
