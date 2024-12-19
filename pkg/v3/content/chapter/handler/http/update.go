package http

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdateChapter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.ChapterCreateUpdateRequest
		var chapter *models.Chapter

		// Parse and validate the request json body
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		// Validate the request validateBody using a validator and return the translated validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Convert the request data to JSON data
		data, err := json.Marshal(requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Convert JSON data to a models.Chapter instance
		err = json.Unmarshal(data, &chapter)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Convert and set the ID from the route parameter to requestBody
		chapter.ID = uint(utils.StringToUint(c.Params("id")))

		// Invoke request to the chapter  usecase to update the chapter
		errMap = handler.usecase.UpdateChapter(*chapter)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
