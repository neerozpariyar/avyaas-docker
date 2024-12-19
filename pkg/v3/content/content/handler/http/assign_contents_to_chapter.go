package http

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) AssignContentToChapterHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.ChapterContent
		errMap := make(map[string]string)

		// Call the repository function to assign content to chapter
		err := c.BodyParser(&req)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		err = validate.Struct(req)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		if err := handler.usecase.AssignContentsToChapter(req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Return success response
		return c.JSON(fiber.Map{"success": true, "message": "Content assigned to chapter successfully"})
	}
}
