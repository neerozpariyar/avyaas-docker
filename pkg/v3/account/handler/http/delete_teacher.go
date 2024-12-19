package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
DeleteTeacher is a Fiber handler function for processing requests to delete a teacher.

Parameters:
  - handler: A pointer to the handler struct, typically representing the application's request
    handler. It is used to access the associated use case for deleting a teacher.

Returns:
  - fiber.Handler: A Fiber handler function that can be registered to handle HTTP requests. The
    handler function processes requests, extracts necessary information, and delegates the teacher
    deletion operation to the use case.
*/
func (handler *handler) DeleteTeacher() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Convert and set the ID from the route parameter
		id := utils.StringToUint(c.Params("id"))

		// Invoke the teacher delete request to the teacher usecase
		err := handler.usecase.DeleteTeacher(uint(id))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
