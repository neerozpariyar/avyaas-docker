package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
DeleteCourseGroup is a Fiber handler function that handles the HTTP DELETE request for deleting a
course group. It parses the course group ID from the request parameters, sends a request to delete
the course group, and responds with appropriate success or error messages in JSON format.
*/
func (handler *handler) DeleteCourseGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Convert and set the ID from the route parameter
		id := utils.StringToUint(c.Params("id"))

		// Invoke the course group update request to the course group usecase
		err := handler.usecase.DeleteCourseGroup(uint(id))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
