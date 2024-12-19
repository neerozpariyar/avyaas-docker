package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) EnrollUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("requester").(uint)
		courseID := uint(utils.StringToUint(c.Params("id")))

		errMap := handler.usecase.EnrollInCourse(userID, courseID)

		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
