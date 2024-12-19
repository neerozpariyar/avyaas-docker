package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdateAvailability() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the test id from params
		courseID := uint(utils.StringToUint(c.Params("id")))

		// Invoke request to the test usecase to update the test status
		if errMap := handler.usecase.UpdateAvailability(courseID); len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
