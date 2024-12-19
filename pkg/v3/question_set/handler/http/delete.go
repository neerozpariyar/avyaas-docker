package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
DeleteQuestionSet is a Fiber handler function that handles the HTTP DELETE request for deleting a
question set. It parses the question set ID from the request parameters, sends a request to delete
the question set, and responds with appropriate success or error messages in JSON format.
*/
func (handler *handler) DeleteQuestionSet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Convert and set the ID from the route parameter
		id := utils.StringToUint(c.Params("id"))

		// Invoke the test delete request to the question set usecase
		if err := handler.usecase.DeleteQuestionSet(uint(id)); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
