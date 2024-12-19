package http

// import (
// 	"avyaas/internal/domain/presenter"
// 	"avyaas/utils"
// 	"net/http"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gofiber/fiber/v2"
// )

// func (handler *handler) UpdateBookmark() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		errMap := make(map[string]string)
// 		var data presenter.BookmarkCreateUpdateRequest
// 		data.UserID = c.Locals("requester").(uint)

// 		// Parse and validate the request json body
// 		err := c.BodyParser(&data)
// 		if err != nil {
// 			errMap["error"] = err.Error()
// 			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 		}

// 		// Initialize the translation function
// 		validate, trans := utils.InitTranslator()

// 		// Validate the request validateBody using a validator and return the translated validation errors if present
// 		err = validate.Struct(data)
// 		if err != nil {
// 			validationErrors := err.(validator.ValidationErrors)
// 			errMap = utils.TranslateError(validationErrors, trans)

// 			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 		}

// 		// Convert and set the ID from the route parameter to requestBody
// 		data.ID = uint(utils.StringToUint(c.Params("id")))

// 		// Invoke request to the bookmark  usecase to update the bookmark
// 		errMap = handler.usecase.UpdateBookmark(data)
// 		if len(errMap) > 0 {
// 			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 		}

// 		return c.JSON(presenter.SuccessResponse())
// 	}
// }