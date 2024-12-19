package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateCourse() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.CourseCreateUpdateRequest
		// var course *models.Course

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

		// Change the CourseID to uppercase
		requestBody.CourseID = strings.ToUpper(requestBody.CourseID)

		// Invoke request to the course usecase to create the course
		errMap = handler.usecase.CreateCourse(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
