package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
DeleteQuestion is a Fiber handler function that handles the HTTP DELETE request for deleting a
question. It parses the question ID from the request parameters, sends a request to delete the
question, and responds with appropriate success or error messages in JSON format.
*/
func (handler *handler) DeleteTypeQuestion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Convert and set the ID from the route parameter
		id := utils.StringToUint(c.Params("id"))

		// Invoke the question delete request to the question usecase
		err := handler.usecase.DeleteTypeQuestion(uint(id))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
