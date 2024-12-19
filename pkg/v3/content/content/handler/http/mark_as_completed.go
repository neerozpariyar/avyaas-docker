package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) MarkAsCompleted() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("requester").(uint)
		contentID := uint(utils.StringToUint(c.Params("id")))

		errMap := handler.usecase.MarkAsCompleted(userID, contentID)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
