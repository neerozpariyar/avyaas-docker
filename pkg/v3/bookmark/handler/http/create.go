package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateBookmark() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var data presenter.BookmarkCreateUpdateRequest
		// data := &presenter.BookmarkCreateUpdateRequest{
		// 	QuestionID: uint(utils.StringToUint(c.Params("questionID"))),
		// 	ContentID:  uint(utils.StringToUint(c.Params("contentID"))),
		// }
		data.UserID = c.Locals("requester").(uint)

		err := c.BodyParser(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Check if the contentID or questionSetID is provided
		if data.ContentID == 0 && data.QuestionID == 0 {
			errMap["error"] = "Either contentID or questionID must be provided"
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		validate, trans := utils.InitTranslator()

		err = validate.Struct(data)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Invoke request to the bookmark usecase to create the bookmark
		errMap = handler.usecase.CreateBookmark(data)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
