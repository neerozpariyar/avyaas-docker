package http

// import (
// 	"avyaas/internal/domain/presenter"
// 	"avyaas/utils"
// 	"net/http"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gofiber/fiber/v2"
// )

// func (handler *handler) SubscribeCourse() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		errMap := make(map[string]string)
// 		var requestBody presenter.SubscribeCourseRequest
// 		requestBody.UserID = c.Locals("requester").(uint)
// 		requestBody.CourseID = uint(utils.StringToUint(c.Params("id")))

// 		err := c.BodyParser(&requestBody)
// 		if err != nil {
// 			errMap["error"] = err.Error()
// 			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 		}

// 		validate, trans := utils.InitTranslator()

// 		err = validate.Struct(requestBody)
// 		if err != nil {
// 			validationErrors := err.(validator.ValidationErrors)
// 			errMap = utils.TranslateError(validationErrors, trans)

// 			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 		}

// 		// Invoke request to the course usecase to create the course
// 		errMap = handler.usecase.SubscribeCourse(requestBody)

// 		if len(errMap) > 0 {
// 			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
// 		}

// 		return c.JSON(presenter.SuccessResponse())
// 	}
// }
