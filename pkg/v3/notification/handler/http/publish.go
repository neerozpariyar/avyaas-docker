package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) PublishNotification() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		errMap := make(map[string]string)

		notifID := uint(utils.StringToUint(c.Params("id")))

		// Invoke request to the notification usecase to create the notification
		err = handler.usecase.PublishNotification(notifID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
