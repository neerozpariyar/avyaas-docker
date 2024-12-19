package handler

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateBlog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.BlogCreateUpdatePresenter

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
		file, _ := c.FormFile("cover")
		if file != nil {
			requestBody.Cover = file
		}

		requestBody.CreatedBy = c.Locals("requester").(uint)

		// Invoke request to the note usecase to create the note
		_, err = handler.usecase.CreateBlog(requestBody)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
