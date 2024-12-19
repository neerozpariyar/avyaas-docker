package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) EvaluateContentProgress() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.ContentProgressPresenter

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
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))
		userID := c.Locals("requester").(uint)

		// Get the course ID associated with the content ID
		courseID, err := handler.usecase.GetCourseIDByContentID(requestBody.ID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Invoke usecase to evaluate content progress
		err = handler.usecase.EvaluateContentProgress(presenter.ContentProgressPresenter{
			UserID: userID,
			// ContentID:       requestBody.ContentID,
			ElapsedDuration: requestBody.ElapsedDuration,
		})

		if err != nil {
			log.Println("Error in EvaluateContentProgress:", err)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		// Invoke usecase to evaluate overall progress
		progressRequest := presenter.ProgressPresenter{
			UserID:   userID,
			CourseID: courseID,
		}
		err = handler.usecase.EvaluateProgress(progressRequest)
		if err != nil {
			log.Println("Error in EvaluateProgress:", err)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		return c.JSON(presenter.SuccessResponse())

		// response := presenter.ListResponse{
		// 	Success: true,
		// 	Data:    nil,
		// }

		// return c.JSON(response)
	}
}
